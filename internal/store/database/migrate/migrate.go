// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package migrate

import (
	"github.com/tphoney/musicscan/internal/store/database/migrate/postgres"
	"github.com/tphoney/musicscan/internal/store/database/migrate/sqlite"

	"github.com/jmoiron/sqlx"
)

// Migrate performs the database migration.
func Migrate(db *sqlx.DB) error {
	switch db.DriverName() {
	case "postgres":
		return postgres.Migrate(db.DB)
	default:
		return sqlite.Migrate(db.DB)
	}
}
