// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package member

import (
	"os"
	"text/template"

	"github.com/tphoney/musicscan/cli/util"
	"github.com/tphoney/musicscan/types"
	"github.com/tphoney/musicscan/types/enum"

	"github.com/alecthomas/kingpin/v2"
	"github.com/drone/funcmap"
)

const ADMIN = "admin"
const DEVELOPER = "developer"

type createCommand struct {
	proj int64
	user string
	role string
	tmpl string
}

func (c *createCommand) run(*kingpin.ParseContext) error {
	client, err := util.Client()
	if err != nil {
		return err
	}
	in := new(types.MembershipInput)
	in.Project = c.proj
	in.User = c.user
	switch c.role {
	case ADMIN:
		in.Role = enum.RoleAdmin
	case DEVELOPER:
		in.Role = enum.RoleDeveloper
	}

	member, err := client.MemberCreate(in)
	if err != nil {
		return err
	}
	tmpl, err := template.New("_").Funcs(funcmap.Funcs).Parse(c.tmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, member)
}

// helper function registers the user create command
func registerCreate(app *kingpin.CmdClause) {
	c := new(createCommand)

	cmd := app.Command("create", "create a project").
		Action(c.run)

	cmd.Arg("project", "project id").
		Required().
		Int64Var(&c.proj)

	cmd.Arg("user id or email", "member id or email").
		Required().
		StringVar(&c.user)

	cmd.Flag("role", "update member role").
		StringVar(&c.role)

	cmd.Flag("format", "format the output using a Go template").
		Default(memberTmpl).
		Hidden().
		StringVar(&c.tmpl)
}
