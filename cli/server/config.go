// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package server

import (
	"os"

	"github.com/tphoney/musicscan/types"

	"github.com/kelseyhightower/envconfig"
)

// legacy environment variables. the key is the legacy
// variable name, and the value is the new variable name.
var legacy = map[string]string{
	// none defined
}

// load returns the system configuration from the
// host environment.
func load() (*types.Config, error) {
	// loop through legacy environment variable and, if set
	// rewrite to the new variable name.
	for k, v := range legacy {
		if s, ok := os.LookupEnv(k); ok {
			os.Setenv(v, s)
		}
	}

	config := new(types.Config)
	// read the configuration from the environment and
	// populate the configuration structure.
	err := envconfig.Process("", config)
	return config, err
}
