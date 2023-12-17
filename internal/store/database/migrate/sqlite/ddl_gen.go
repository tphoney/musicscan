package sqlite

import (
	"database/sql"
)

var migrations = []struct {
	name string
	stmt string
}{
	{
		name: "create-table-users",
		stmt: createTableUsers,
	},
	{
		name: "create-table-projects",
		stmt: createTableProjects,
	},
	{
		name: "create-table-members",
		stmt: createTableMembers,
	},
	{
		name: "create-index-members-project-id",
		stmt: createIndexMembersProjectId,
	},
	{
		name: "create-index-members-user-id",
		stmt: createIndexMembersUserId,
	},
	{
		name: "create-table-artists",
		stmt: createTableArtists,
	},
	{
		name: "create-index-artist-project-id",
		stmt: createIndexArtistProjectId,
	},
	{
		name: "create-table-albums",
		stmt: createTableAlbums,
	},
	{
		name: "create-index-album-artist-id",
		stmt: createIndexAlbumArtistId,
	},
}

// Migrate performs the database migration. If the migration fails
// and error is returned.
func Migrate(db *sql.DB) error {
	if err := createTable(db); err != nil {
		return err
	}
	completed, err := selectCompleted(db)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	for _, migration := range migrations {
		if _, ok := completed[migration.name]; ok {

			continue
		}

		if _, err := db.Exec(migration.stmt); err != nil {
			return err
		}
		if err := insertMigration(db, migration.name); err != nil {
			return err
		}

	}
	return nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(migrationTableCreate)
	return err
}

func insertMigration(db *sql.DB, name string) error {
	_, err := db.Exec(migrationInsert, name)
	return err
}

func selectCompleted(db *sql.DB) (map[string]struct{}, error) {
	migrations := map[string]struct{}{}
	rows, err := db.Query(migrationSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		migrations[name] = struct{}{}
	}
	return migrations, nil
}

//
// migration table ddl and sql
//

var migrationTableCreate = `
CREATE TABLE IF NOT EXISTS migrations (
 name VARCHAR(255)
,UNIQUE(name)
)
`

var migrationInsert = `
INSERT INTO migrations (name) VALUES (?)
`

var migrationSelect = `
SELECT name FROM migrations
`

//
// 001_create_table_user.sql
//

var createTableUsers = `
CREATE TABLE IF NOT EXISTS users (
 user_id            INTEGER PRIMARY KEY AUTOINCREMENT
,user_email         TEXT COLLATE NOCASE
,user_password      TEXT
,user_token         TEXT
,user_name          TEXT
,user_company       TEXT
,user_admin         BOOLEAN
,user_blocked       BOOLEAN
,user_created       INTEGER
,user_updated       INTEGER
,user_authed        INTEGER
,UNIQUE(user_token)
,UNIQUE(user_email COLLATE NOCASE)
);
`

//
// 002_create_table_project.sql
//

var createTableProjects = `
CREATE TABLE IF NOT EXISTS projects (
 project_id          INTEGER PRIMARY KEY AUTOINCREMENT
,project_name        TEXT
,project_desc        TEXT
,project_token       TEXT
,project_active      BOOLEAN
,project_created     INTEGER
,project_updated     INTEGER
,UNIQUE(project_token)
);
`

//
// 003_create_table_member.sql
//

var createTableMembers = `
CREATE TABLE IF NOT EXISTS members (
 member_project_id INTEGER
,member_user_id    INTEGER
,member_role       INTEGER
,PRIMARY KEY(member_project_id, member_user_id)
);
`

var createIndexMembersProjectId = `
CREATE INDEX IF NOT EXISTS index_members_project ON members(member_project_id)
`

var createIndexMembersUserId = `
CREATE INDEX IF NOT EXISTS index_members_user ON members(member_user_id)
`

//
// 005_create_table_artist.sql
//

var createTableArtists = `
CREATE TABLE IF NOT EXISTS artists (
 artist_id          INTEGER PRIMARY KEY AUTOINCREMENT
,artist_project_id  INTEGER
,artist_name        TEXT
,artist_desc        TEXT
,artist_wanted      BOOLEAN DEFAULT 1
,artist_popularity  INTEGER DEFAULT 0
,artist_spotify     TEXT
,artist_created     INTEGER
,artist_updated     INTEGER
);
`

var createIndexArtistProjectId = `
CREATE INDEX IF NOT EXISTS index_artist_project ON artists(artist_project_id);
`

//
// 007_create_table_album.sql
//

var createTableAlbums = `
CREATE TABLE IF NOT EXISTS albums (
 album_id        INTEGER PRIMARY KEY AUTOINCREMENT
,album_artist_id INTEGER
,album_name      TEXT
,album_desc      TEXT
,album_format    TEXT
,album_wanted    BOOLEAN DEFAULT 1
,album_year      TEXT
,album_created   INTEGER
,album_updated   INTEGER
);
`

var createIndexAlbumArtistId = `
CREATE INDEX IF NOT EXISTS index_album_artist ON albums(album_artist_id);
`
