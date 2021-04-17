// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package cli

import (
	"os"

	"github.com/tphoney/musicscan/cli/album"
	"github.com/tphoney/musicscan/cli/artist"
	"github.com/tphoney/musicscan/cli/member"
	"github.com/tphoney/musicscan/cli/project"
	"github.com/tphoney/musicscan/cli/server"
	"github.com/tphoney/musicscan/cli/token"
	"github.com/tphoney/musicscan/cli/user"
	"github.com/tphoney/musicscan/cli/users"
	"github.com/tphoney/musicscan/version"

	"gopkg.in/alecthomas/kingpin.v2"
)

// empty context
// var nocontext = context.Background()

// application name
var application = "musicscan"

// application description
var description = "description goes here" // TODO edit this application description

// Command parses the command line arguments and then executes a
// subcommand program.
func Command() {
	app := kingpin.New(application, description)
	server.Register(app)
	user.Register(app)
	project.Register(app)
	artist.Register(app)
	album.Register(app)
	member.Register(app)
	users.Register(app)
	token.Register(app)
	registerLogin(app)
	registerLogout(app)
	registerRegister(app)
	registerScan(app)

	kingpin.Version(version.Version.String())
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
