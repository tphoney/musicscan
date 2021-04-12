// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package dist embeds the static web server content.
package dist

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed files/*
var content embed.FS

// FileSystem provides access to the static web server
// content, embedded in the binary.
func FileSystem() http.FileSystem {
	fsys, _ := fs.Sub(content, "files")
	return http.FS(fsys)
}
