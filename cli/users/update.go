// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package users

import (
	"os"
	"text/template"

	"github.com/tphoney/musicscan/cli/util"
	"github.com/tphoney/musicscan/types"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/guregu/null.v4"
)

type updateCommand struct {
	id     string
	email  string
	admin  bool
	demote bool
	tmpl   string
}

func (c *updateCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}

	in := new(types.UserInput)
	if v := c.email; v != "" {
		in.Username = null.StringFrom(v)
	}
	if v := c.admin; v {
		in.Admin = null.BoolFrom(v)
	}
	if v := c.demote; v {
		in.Admin = null.BoolFrom(false)
	}

	user, err := client.UserUpdate(c.id, in)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, user)
}

// helper function registers the user update command
func registerUpdate(app *kingpin.CmdClause) {
	c := new(updateCommand)

	cmd := app.Command("update", "update a user").
		Action(c.run)

	cmd.Arg("id or email", "user id or email").
		Required().
		StringVar(&c.id)

	cmd.Flag("email", "update user email").
		StringVar(&c.email)

	cmd.Flag("promote", "promote user to admin").
		BoolVar(&c.admin)

	cmd.Flag("demote", "demote user from admin").
		BoolVar(&c.demote)

	cmd.Flag("format", "format the output using a Go template").
		Default(userTmpl).
		Hidden().
		StringVar(&c.tmpl)
}
