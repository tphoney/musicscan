// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package project

import (
	"github.com/tphoney/musicscan/cli/util"

	"github.com/alecthomas/kingpin/v2"
)

type deleteCommand struct {
	id int64
}

func (c *deleteCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	return client.ProjectDelete(c.id)
}

// helper function registers the user delete command
func registerDelete(app *kingpin.CmdClause) {
	c := new(deleteCommand)

	cmd := app.Command("delete", "delete a project").
		Action(c.run)

	cmd.Arg("id ", "project id").
		Required().
		Int64Var(&c.id)
}
