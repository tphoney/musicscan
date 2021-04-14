// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package users

import (
	"os"
	"text/template"

	"github.com/drone/funcmap"
	"github.com/tphoney/musicscan/cli/util"

	"gopkg.in/alecthomas/kingpin.v2"
)

const userTmpl = `
id:    {{ .ID }}
email: {{ .Email }}
admin: {{ .Admin }}
`

type listCommand struct {
	tmpl string
}

func (c *listCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	list, err := client.UserList()
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl + "\n")
	if err != nil {
		return err
	}
	for _, item := range list {
		tmpl.Execute(os.Stdout, item)
	}
	return nil
}

// helper function registers the user list command
func registerList(app *kingpin.CmdClause) {
	c := new(listCommand)

	cmd := app.Command("ls", "display a list of users").
		Action(c.run)

	cmd.Flag("format", "format the output using a Go template").
		Default(userTmpl).
		Hidden().
		StringVar(&c.tmpl)
}
