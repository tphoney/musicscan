// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package album

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
	proj int64
	artist  int64
	album  int64
	name string
	desc string
	tmpl string
}

func (c *updateCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}

	in := new(types.AlbumInput)
	if v := c.name; v != "" {
		in.Name = null.StringFrom(v)
	}
	if v := c.desc; v != "" {
		in.Desc = null.StringFrom(v)
	}

	artist, err := client.AlbumUpdate(c.proj, c.artist, c.album, in)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, artist)
}

// helper function registers the update command
func registerUpdate(app *kingpin.CmdClause) {
	c := new(updateCommand)

	cmd := app.Command("update", "update a artist").
		Action(c.run)

	cmd.Arg("project_id", "project id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("artist_id", "artist id").
		Required().
		Int64Var(&c.artist)

	cmd.Arg("album_id", "album id").
		Required().
		Int64Var(&c.album)

	cmd.Flag("name", "update project name").
		StringVar(&c.name)

	cmd.Flag("desc", "update project description").
		StringVar(&c.desc)

	cmd.Flag("format", "format the output using a Go template").
		Default(projectTmpl).
		Hidden().
		StringVar(&c.tmpl)
}
