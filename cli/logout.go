// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package cli

import (
	"os"

	"github.com/tphoney/musicscan/cli/util"

	"github.com/alecthomas/kingpin/v2"
)

type logoutCommand struct{}

func (c *logoutCommand) run(*kingpin.ParseContext) error {
	path, err := util.Config()
	if err != nil {
		return err
	}
	return os.Remove(path)
}

// helper function to register the logout command.
func registerLogout(app *kingpin.Application) {
	c := new(logoutCommand)

	app.Command("logout", "logout from the remote server").
		Action(c.run)
}
