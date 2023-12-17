// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// +build proxy

package dist

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Handler returns an http.HandlerFunc that servers the
// static content from the embedded file system.
func Handler() http.HandlerFunc {
	downstream, _ := url.Parse("http://localhost:8080")
	return httputil.NewSingleHostReverseProxy(downstream).ServeHTTP
}
