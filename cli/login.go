// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package cli

import (
	"encoding/json"
	"io/ioutil"

	"github.com/tphoney/musicscan/cli/util"
	"github.com/tphoney/musicscan/client"

	"gopkg.in/alecthomas/kingpin.v2"
)

type loginCommand struct {
	server string
}

func (c *loginCommand) run(*kingpin.ParseContext) error {
	username, password := util.Credentials()
	client := client.New(c.server)
	token, err := client.Login(username, password)
	if err != nil {
		return err
	}
	path, err := util.Config()
	if err != nil {
		return err
	}
	token.Address = c.server
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0600)
}

// helper function to register the logout command.
func registerLogin(app *kingpin.Application) {
	c := new(loginCommand)

	cmd := app.Command("login", "login to the remote server").
		Action(c.run)

	cmd.Arg("server", "server address").
		Default("http://localhost:3000").
		StringVar(&c.server)
}
