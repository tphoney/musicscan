// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package system

import (
	"fmt"
	"net/http"

	"github.com/tphoney/musicscan/version"
)

// HandleVersion writes the server version number
// to the http.Response body in plain text.
func HandleVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", version.Version)
}
