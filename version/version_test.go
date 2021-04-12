// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package version

import "testing"

func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.0.0"; got != want {
		t.Errorf("Want version %s, got %s", want, got)
	}
}
