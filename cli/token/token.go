// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package token

import (
	"github.com/tphoney/musicscan/cli/util"

	"github.com/alecthomas/kingpin/v2"
)

type command struct {
}

func (c *command) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	token, err := client.Token()
	if err != nil {
		return err
	}
	println(token.Value)
	return nil
}

// Register the command.
func Register(app *kingpin.Application) {
	c := new(command)

	app.Command("token", "generate a personal token").
		Action(c.run)
}
