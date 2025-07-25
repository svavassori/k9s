// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package view

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/derailed/k9s/internal"
	"github.com/derailed/k9s/internal/client"
	"github.com/derailed/k9s/internal/dao"
	"github.com/derailed/k9s/internal/port"
	"github.com/derailed/k9s/internal/slogs"
	"github.com/derailed/k9s/internal/ui"
	"github.com/derailed/k9s/internal/watch"
	"github.com/derailed/tcell/v2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/portforward"
)

// PortForwardExtender adds port-forward extensions.
type PortForwardExtender struct {
	ResourceViewer
}

// NewPortForwardExtender returns a new extender.
func NewPortForwardExtender(r ResourceViewer) ResourceViewer {
	p := PortForwardExtender{ResourceViewer: r}
	p.AddBindKeysFn(p.bindKeys)

	return &p
}

func (p *PortForwardExtender) bindKeys(aa *ui.KeyActions) {
	aa.Bulk(ui.KeyMap{
		ui.KeyF:      ui.NewKeyAction("Show PortForward", p.showPFCmd, true),
		ui.KeyShiftF: ui.NewKeyAction("Port-Forward", p.portFwdCmd, true),
	})
}

func (p *PortForwardExtender) portFwdCmd(evt *tcell.EventKey) *tcell.EventKey {
	path := p.GetTable().GetSelectedItem()
	if path == "" {
		return evt
	}

	podName, err := p.fetchPodName(path)
	if err != nil {
		p.App().Flash().Err(err)
		return nil
	}
	if err := ensurePodPortFwdAllowed(p.App().factory, podName); err != nil {
		p.App().Flash().Err(err)
		return nil
	}
	if err := showFwdDialog(p, podName, startFwdCB); err != nil {
		p.App().Flash().Err(err)
	}

	return nil
}

func (p *PortForwardExtender) showPFCmd(evt *tcell.EventKey) *tcell.EventKey {
	path := p.GetTable().GetSelectedItem()
	if path == "" {
		return evt
	}

	podName, err := p.fetchPodName(path)
	if err != nil {
		p.App().Flash().Err(err)
		return nil
	}

	if !p.App().factory.Forwarders().IsPodForwarded(podName) {
		p.App().Flash().Errf("no port-forward defined")
		return nil
	}

	pf := NewPortForward(client.PfGVR)
	pf.SetContextFn(p.portForwardContext)
	if err := p.App().inject(pf, false); err != nil {
		p.App().Flash().Err(err)
	}

	return nil
}

func (p *PortForwardExtender) fetchPodName(path string) (string, error) {
	res, err := dao.AccessorFor(p.App().factory, p.GVR())
	if err != nil {
		return "", err
	}
	ctrl, ok := res.(dao.Controller)
	if !ok {
		return "", fmt.Errorf("expecting a controller resource for %q", p.GVR())
	}

	return ctrl.Pod(path)
}

func (p *PortForwardExtender) portForwardContext(ctx context.Context) context.Context {
	if bc := p.App().BenchFile; bc != "" {
		ctx = context.WithValue(ctx, internal.KeyBenchCfg, p.App().BenchFile)
	}

	return context.WithValue(ctx, internal.KeyPath, p.GetTable().GetSelectedItem())
}

// ----------------------------------------------------------------------------
// Helpers...

func ensurePodPortFwdAllowed(factory dao.Factory, podName string) error {
	pod, err := fetchPod(factory, podName)
	if err != nil {
		return err
	}
	if pod.Status.Phase != v1.PodRunning {
		return fmt.Errorf("pod must be running. Current status=%v", pod.Status.Phase)
	}

	return nil
}

func runForward(v ResourceViewer, pf watch.Forwarder, f *portforward.PortForwarder) {
	v.App().factory.AddForwarder(pf)

	v.App().QueueUpdateDraw(func() {
		DismissPortForwards(v, v.App().Content.Pages)
	})

	pf.SetActive(true)
	if err := f.ForwardPorts(); err != nil {
		v.App().Flash().Warnf("PortForward failed for %s: %s. Deleting!", pf.ID(), err)
	}
	v.App().QueueUpdateDraw(func() {
		v.App().factory.DeleteForwarder(pf.ID())
		pf.SetActive(false)
	})
}

func startFwdCB(v ResourceViewer, path string, pts port.PortTunnels) error {
	if err := pts.CheckAvailable(); err != nil {
		return err
	}

	tt := make([]string, 0, len(pts))
	for _, pt := range pts {
		if _, ok := v.App().factory.ForwarderFor(dao.PortForwardID(path, pt.Container, pt.PortMap())); ok {
			return fmt.Errorf("port-forward is already active on pod %s", path)
		}
		pf := dao.NewPortForwarder(v.App().factory)
		fwd, err := pf.Start(path, pt)
		if err != nil {
			return err
		}
		slog.Debug(">>> Starting port forward",
			slogs.PFID, pf.ID(),
			slogs.PFTunnel, pt,
		)
		go runForward(v, pf, fwd)
		tt = append(tt, pt.LocalPort)
	}
	if len(tt) == 1 {
		v.App().Flash().Infof("PortForward activated %s", tt[0])
		return nil
	}
	v.App().Flash().Infof("PortForwards activated %s", strings.Join(tt, ","))

	return nil
}

func showFwdDialog(v ResourceViewer, path string, cb PortForwardCB) error {
	mm, anns, err := fetchPodPorts(v.App().factory, path)
	if err != nil {
		return err
	}
	ports := make(port.ContainerPortSpecs, 0, len(mm))
	for co, pp := range mm {
		for _, p := range pp {
			if p.Protocol != v1.ProtocolTCP {
				continue
			}
			ports = append(ports, port.NewPortSpec(co, p.Name, p.ContainerPort))
		}
	}
	if spec, ok := anns[port.K9sAutoPortForwardsKey]; ok {
		pfs, err := port.ParsePFs(spec)
		if err != nil {
			return err
		}

		pts, err := pfs.ToTunnels(v.App().Config.K9s.PortForwardAddress, ports, port.IsPortFree)
		if err != nil {
			return err
		}

		return startFwdCB(v, path, pts)
	}
	ShowPortForwards(v, path, ports, anns, cb)

	return nil
}

func fetchPodPorts(f *watch.Factory, path string) (ports map[string][]v1.ContainerPort, anns map[string]string, err error) {
	slog.Debug("Fetching ports on pod", slogs.FQN, path)
	o, err := f.Get(client.PodGVR, path, true, labels.Everything())
	if err != nil {
		return nil, nil, err
	}

	var pod v1.Pod
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.(*unstructured.Unstructured).Object, &pod)
	if err != nil {
		return nil, nil, err
	}

	pp := make(map[string][]v1.ContainerPort, len(pod.Spec.Containers))
	for i := range pod.Spec.Containers {
		pp[pod.Spec.Containers[i].Name] = pod.Spec.Containers[i].Ports
	}

	return pp, pod.Annotations, nil
}
