// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package server

import (
	"net/http"

	"github.com/tphoney/musicscan/types"

	"github.com/google/wire"
)

// WireSet provides a wire set for this package
var WireSet = wire.NewSet(ProvideServer)

// ProvideServer provides a server instance
func ProvideServer(config *types.Config, handler http.Handler) *Server {
	return &Server{
		Acme:    config.Server.Acme.Enabled,
		Addr:    config.Server.Bind,
		Host:    config.Server.Host,
		Handler: handler,
	}
}
