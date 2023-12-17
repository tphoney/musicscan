// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package user

import (
	"os"
	"text/template"

	"github.com/tphoney/musicscan/cli/util"

	"github.com/alecthomas/kingpin/v2"
	"github.com/drone/funcmap"
)

const userTmpl = `
email: {{ .Email }}
admin: {{ .Admin }}
`

type command struct {
	tmpl string
}

func (c *command) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	user, err := client.Self()
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, user)
}

// Register the command.
func Register(app *kingpin.Application) {
	c := new(command)

	cmd := app.Command("account", "display authenticated user").
		Action(c.run)

	cmd.Flag("format", "format the output using a Go template").
		Default(userTmpl).
		Hidden().
		StringVar(&c.tmpl)
}
