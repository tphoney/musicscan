// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package util

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/adrg/xdg"
	"github.com/tphoney/musicscan/client"
	"github.com/tphoney/musicscan/types"
	"golang.org/x/term"
)

// Client returns a client that is configured from file.
func Client() (*client.HTTPClient, error) {
	path, err := Config()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	token := new(types.Token)
	if err := json.Unmarshal(data, token); err != nil {
		return nil, err
	}
	if time.Now().Unix() > token.Expires.Unix() {
		return nil, errors.New("token is expired, please login")
	}
	cli := client.NewToken(token.Address, token.Value)
	if os.Getenv("DEBUG") == "true" {
		cli.SetDebug(true)
	}
	return cli, nil
}

// Config returns the configuration file path.
func Config() (string, error) {
	return xdg.ConfigFile(
		filepath.Join("app", "config.json"),
	)
}

// Credentials returns the username and password from stdin.
func Credentials() (username, password string) {
	return Username(), Password()
}

// Username returns the username from stdin.
func Username() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	return strings.TrimSpace(username)
}

// Password returns the password from stdin.
func Password() string {
	fmt.Print("Enter Password: ")
	passwordb, _ := term.ReadPassword(syscall.Stdin)
	password := string(passwordb)

	return strings.TrimSpace(password)
}
