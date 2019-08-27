// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/derailed/k9s/internal/resource (interfaces: Cruder)

package resource_test

import (
	k8s "github.com/derailed/k9s/internal/k8s"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
	"time"
)

type MockCruder struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCruder(options ...pegomock.Option) *MockCruder {
	mock := &MockCruder{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockCruder) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockCruder) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockCruder) Delete(_param0 string, _param1 string, _param2 bool, _param3 bool) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCruder().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Delete", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCruder) Get(_param0 string, _param1 string) (interface{}, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCruder().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Get", params, []reflect.Type{reflect.TypeOf((*interface{})(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 interface{}
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(interface{})
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCruder) GetFieldSelector() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCruder().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetFieldSelector", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockCruder) GetLabelSelector() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCruder().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetLabelSelector", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockCruder) HasSelectors() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCruder().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("HasSelectors", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockCruder) List(_param0 string) (k8s.Collection, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCruder().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("List", params, []reflect.Type{reflect.TypeOf((*k8s.Collection)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 k8s.Collection
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(k8s.Collection)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCruder) SetFieldSelector(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCruder().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetFieldSelector", params, []reflect.Type{})
}

func (mock *MockCruder) SetLabelSelector(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCruder().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetLabelSelector", params, []reflect.Type{})
}

func (mock *MockCruder) VerifyWasCalledOnce() *VerifierMockCruder {
	return &VerifierMockCruder{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockCruder) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockCruder {
	return &VerifierMockCruder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockCruder) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockCruder {
	return &VerifierMockCruder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockCruder) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockCruder {
	return &VerifierMockCruder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockCruder struct {
	mock                   *MockCruder
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockCruder) Delete(_param0 string, _param1 string, _param2 bool, _param3 bool) *MockCruder_Delete_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Delete", params, verifier.timeout)
	return &MockCruder_Delete_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCruder_Delete_OngoingVerification struct {
	mock              *MockCruder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCruder_Delete_OngoingVerification) GetCapturedArguments() (string, string, bool, bool) {
	_param0, _param1, _param2, _param3 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1]
}

func (c *MockCruder_Delete_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []bool, _param3 []bool) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]bool, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(bool)
		}
		_param3 = make([]bool, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(bool)
		}
	}
	return
}

func (verifier *VerifierMockCruder) Get(_param0 string, _param1 string) *MockCruder_Get_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Get", params, verifier.timeout)
	return &MockCruder_Get_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCruder_Get_OngoingVerification struct {
	mock              *MockCruder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCruder_Get_OngoingVerification) GetCapturedArguments() (string, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockCruder_Get_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockCruder) GetFieldSelector() *MockCruder_GetFieldSelector_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetFieldSelector", params, verifier.timeout)
	return &MockCruder_GetFieldSelector_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCruder_GetFieldSelector_OngoingVerification struct {
	mock              *MockCruder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCruder_GetFieldSelector_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCruder_GetFieldSelector_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCruder) GetLabelSelector() *MockCruder_GetLabelSelector_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetLabelSelector", params, verifier.timeout)
	return &MockCruder_GetLabelSelector_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCruder_GetLabelSelector_OngoingVerification struct {
	mock              *MockCruder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCruder_GetLabelSelector_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCruder_GetLabelSelector_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCruder) HasSelectors() *MockCruder_HasSelectors_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "HasSelectors", params, verifier.timeout)
	return &MockCruder_HasSelectors_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCruder_HasSelectors_OngoingVerification struct {
	mock              *MockCruder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCruder_HasSelectors_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCruder_HasSelectors_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCruder) List(_param0 string) *MockCruder_List_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "List", params, verifier.timeout)
	return &MockCruder_List_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCruder_List_OngoingVerification struct {
	mock              *MockCruder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCruder_List_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockCruder_List_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockCruder) SetFieldSelector(_param0 string) *MockCruder_SetFieldSelector_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetFieldSelector", params, verifier.timeout)
	return &MockCruder_SetFieldSelector_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCruder_SetFieldSelector_OngoingVerification struct {
	mock              *MockCruder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCruder_SetFieldSelector_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockCruder_SetFieldSelector_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockCruder) SetLabelSelector(_param0 string) *MockCruder_SetLabelSelector_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetLabelSelector", params, verifier.timeout)
	return &MockCruder_SetLabelSelector_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCruder_SetLabelSelector_OngoingVerification struct {
	mock              *MockCruder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCruder_SetLabelSelector_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockCruder_SetLabelSelector_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}
