// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package artist

import (
	"os"
	"text/template"

	"github.com/tphoney/musicscan/cli/util"
	"github.com/tphoney/musicscan/types"

	"github.com/alecthomas/kingpin/v2"
	"github.com/drone/funcmap"
)

type createCommand struct {
	proj   int64
	name   string
	desc   string
	wanted bool
	tmpl   string
}

func (c *createCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	in := &types.Artist{
		Name:   c.name,
		Desc:   c.desc,
		Wanted: c.wanted,
	}
	proj, err := client.ArtistCreate(c.proj, in)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, proj)
}

// helper function registers the user create command
func registerCreate(app *kingpin.CmdClause) {
	c := new(createCommand)

	cmd := app.Command("create", "create a artist").
		Action(c.run)

	cmd.Arg("project_id", "project id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("name", "artist name").
		Required().
		StringVar(&c.name)

	cmd.Flag("desc", "artist description").
		StringVar(&c.desc)

	cmd.Flag("wanted", "wanted").
		BoolVar(&c.wanted)

	cmd.Flag("format", "format the output using a Go template").
		Default(projectTmpl).
		Hidden().
		StringVar(&c.tmpl)
}
