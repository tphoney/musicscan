// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package artist

import "gopkg.in/alecthomas/kingpin.v2"

// Register the command.
func Register(app *kingpin.Application) {
	cmd := app.Command("artist", "manage artists")
	registerFind(cmd)
	registerList(cmd)
	registerCreate(cmd)
	registerUpdate(cmd)
	registerDelete(cmd)
}
