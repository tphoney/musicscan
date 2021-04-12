// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package router

import "testing"

// this unit test ensures routes that require authorization
// return a 401 unauthorized if no token, or an invalid token
// is provided.
func TestTokenGate(t *testing.T) {
	t.Skip()
}

// this unit test ensures routes that require project access
// return a 403 forbidden if the user does not have acess
// to the project
func TestProjectGate(t *testing.T) {
	t.Skip()
}

// this unit test ensures routes that require system access
// return a 403 forbidden if the user does not have acess
// to the project
func TestSystemGate(t *testing.T) {
	t.Skip()
}
