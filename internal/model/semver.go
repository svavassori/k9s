// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package model

import (
	"fmt"
	"regexp"
	"strconv"
)

var versionRX = regexp.MustCompile(`\Av(\d+)\.(\d+)\.(\d+)(-[a-z]+)?\z`)

// SemVer represents a semantic version.
type SemVer struct {
	Major, Minor, Patch int
	Build               string
}

// NewSemVer returns a new semantic version.
func NewSemVer(version string) *SemVer {
	var v SemVer
	v.Major, v.Minor, v.Patch, v.Build = v.parse(NormalizeVersion(version))

	return &v
}

// String returns version as a string.
func (v *SemVer) String() string {
	return fmt.Sprintf("v%d.%d.%d%s", v.Major, v.Minor, v.Patch, v.Build)
}

func (*SemVer) parse(version string) (major, minor, patch int, build string) {
	mm := versionRX.FindStringSubmatch(version)
	if len(mm) < 5 {
		return
	}
	major, _ = strconv.Atoi(mm[1])
	minor, _ = strconv.Atoi(mm[2])
	patch, _ = strconv.Atoi(mm[3])
	build = mm[4]

	return
}

// NormalizeVersion ensures the version starts with a v.
func NormalizeVersion(version string) string {
	if version == "" {
		return version
	}
	if version[0] == 'v' {
		return version
	}
	return "v" + version
}

// IsCurrent asserts if at latest release.
func (v *SemVer) IsCurrent(latest *SemVer) bool {
	return v.Major >= latest.Major && v.Minor >= latest.Minor && v.Patch >= latest.Patch
}
