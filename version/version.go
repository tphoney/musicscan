// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package version provides the version number.
package version

import "github.com/coreos/go-semver/semver"

var (
	// GitRepository is the git repository that was compiled
	GitRepository string
	// GitCommit is the git commit that was compiled
	GitCommit string
	// VersionMajor is for an API incompatible changes.
	VersionMajor int64 = 1
	// VersionMinor is for functionality in a backwards-compatible manner.
	VersionMinor int64
	// VersionPatch is for backwards-compatible bug fixes.
	VersionPatch int64
	// VersionPre indicates prerelease.
	VersionPre = ""
	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev string
)

// Version is the specification version that the package types support.
var Version = semver.Version{
	Major:      VersionMajor,
	Minor:      VersionMinor,
	Patch:      VersionPatch,
	PreRelease: semver.PreRelease(VersionPre),
	Metadata:   VersionDev,
}
