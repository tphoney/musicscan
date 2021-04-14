// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package artist

import (
	"os"
	"text/template"

	"github.com/drone/funcmap"
	"github.com/tphoney/musicscan/cli/util"

	"gopkg.in/alecthomas/kingpin.v2"
)

const projectTmpl = `
id:   {{ .ID }}
name: {{ .Name }}
desc: {{ .Desc }}
`

type listCommand struct {
	proj int64
	tmpl string
}

func (c *listCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	list, err := client.ArtistList(c.proj)
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

// helper function registers the list command
func registerList(app *kingpin.CmdClause) {
	c := new(listCommand)

	cmd := app.Command("ls", "display a list of artists").
		Action(c.run)

	cmd.Arg("project_id", "project id").
		Required().
		Int64Var(&c.proj)

	cmd.Flag("format", "format the output using a Go template").
		Default(projectTmpl).
		Hidden().
		StringVar(&c.tmpl)
}
