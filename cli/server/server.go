// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package server

import (
	"context"
	"os"

	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/types"
	"github.com/tphoney/musicscan/version"

	"github.com/alecthomas/kingpin/v2"
	"github.com/joho/godotenv"
	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
)

type command struct {
	envfile string
}

func (c *command) run(*kingpin.ParseContext) error {
	// load environment variables from file.
	err := godotenv.Load(c.envfile)
	if err != nil {
		logrus.Warnf("No env %s", err.Error())
	}

	// create the system configuration store by loading
	// data from the environment.
	config, err := load()
	if err != nil {
		logrus.Warnf("No config %s", err.Error())
	}

	// configure the log level
	setupLogger(config)

	server, err := initServer(config)
	if err != nil {
		logrus.Warn(err)
	}

	// create the http server.
	// server := server.Server{
	// 	Acme:    config.Server.Acme.Enabled,
	// 	Addr:    config.Server.Bind,
	// 	Handler: handler,
	// }

	logrus.
		WithField("revision", version.GitCommit).
		WithField("repository", version.GitRepository).
		WithField("version", version.Version).
		Infof("server listening at address %s://%s%s", config.Server.Proto, config.Server.Host, config.Server.Bind)
	// starts the http server.s
	return server.ListenAndServe(context.Background())
}

// helper function configures the global logger from
// the loaded configuration.
func setupLogger(config *types.Config) {
	logger.L = logrus.NewEntry(
		logrus.StandardLogger(),
	)

	// configure the log level
	switch {
	case config.Trace:
		logrus.SetLevel(logrus.TraceLevel)
	case config.Debug:
		logrus.SetLevel(logrus.DebugLevel)
	}

	// if the terminal is not a tty we should output the
	// logs in json format
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

// Register the server command.
func Register(app *kingpin.Application) {
	c := new(command)

	cmd := app.Command("server", "starts the server").
		Action(c.run)

	cmd.Arg("envfile", "load the environment variable file").
		Default("").
		StringVar(&c.envfile)
}
