// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var noContext = context.Background()

// connect opens a new test database connection.
func connect() (*sqlx.DB, error) {
	var (
		driver = "sqlite3"
		config = ":memory:?_foreign_keys=1"
	)
	if os.Getenv("DATABASE_CONFIG") != "" {
		driver = os.Getenv("DATABASE_DRIVER")
		config = os.Getenv("DATABASE_CONFIG")
	}
	return Connect(driver, config)
}

// seed seed the database state.
func seed(db *sqlx.DB) error {
	db.Exec("TRUNCATE TABLE albums")
	db.Exec("TRUNCATE TABLE artists")
	db.Exec("TRUNCATE TABLE members")
	db.Exec("TRUNCATE TABLE projects")
	db.Exec("TRUNCATE TABLE users")

	out, err := ioutil.ReadFile("testdata/seed.sql")
	if err != nil {
		return err
	}
	parts := strings.Split(string(out), ";")
	for _, stmt := range parts {
		if stmt == "" {
			continue
		}
		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("%s: %s", err, stmt)
		}
	}
	return nil
}
