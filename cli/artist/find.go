// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package artist

import (
	"os"
	"text/template"

	"github.com/tphoney/musicscan/cli/util"

	"github.com/drone/funcmap"
	"gopkg.in/alecthomas/kingpin.v2"
)

type findCommand struct {
	proj int64
	id   int64
	tmpl string
}

func (c *findCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	proj, err := client.artist(c.proj, c.id)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl + "\n")
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, proj)
}

// helper function registers the user find command
func registerFind(app *kingpin.CmdClause) {
	c := new(findCommand)

	cmd := app.Command("find", "display project details").
		Action(c.run)

	cmd.Arg("project_id", "project id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("artist_id", "artist id").
		Required().
		Int64Var(&c.id)

	cmd.Flag("format", "format the output using a Go template").
		Default(projectTmpl).
		Hidden().
		StringVar(&c.tmpl)
}
