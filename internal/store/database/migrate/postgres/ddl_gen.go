package postgres

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
		stmt: createTableartists,
	},
	{
		name: "create-index-artist-project-id",
		stmt: createIndexartistProjectId,
	},
	{
		name: "create-table-albums",
		stmt: createTablealbums,
	},
	{
		name: "create-index-album-artist-id",
		stmt: createIndexalbumartistId,
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
INSERT INTO migrations (name) VALUES ($1)
`

var migrationSelect = `
SELECT name FROM migrations
`

//
// 001_create_table_user.sql
//

var createTableUsers = `
CREATE TABLE IF NOT EXISTS users (
 user_id            SERIAL PRIMARY KEY
,user_email         VARCHAR(250)
,user_password      VARCHAR(250)
,user_token         VARCHAR(250)
,user_admin         BOOLEAN
,user_blocked       BOOLEAN
,user_created       INTEGER
,user_updated       INTEGER
,user_authed        INTEGER
,UNIQUE(user_token)
);
`

//
// 002_create_table_project.sql
//

var createTableProjects = `
CREATE TABLE IF NOT EXISTS projects (
 project_id          SERIAL PRIMARY KEY
,project_name        VARCHAR(250)
,project_desc        VARCHAR(250)
,project_token       VARCHAR(250)
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

var createTableartists = `
CREATE TABLE IF NOT EXISTS artists (
 artist_id          SERIAL PRIMARY KEY
,artist_project_id  INTEGER
,artist_name        VARCHAR(250)
,artist_desc        VARCHAR(2000)
,artist_created     INTEGER
,artist_updated     INTEGER
);
`

var createIndexartistProjectId = `
CREATE INDEX IF NOT EXISTS index_artist_project ON artists(artist_project_id);
`

//
// 007_create_table_album.sql
//

var createTablealbums = `
CREATE TABLE IF NOT EXISTS albums (
 album_id       SERIAL PRIMARY KEY
,album_artist_id   INTEGER
,album_name     VARCHAR(250)
,album_desc     VARCHAR(2000)
,album_created  INTEGER
,album_updated  INTEGER
);
`

var createIndexalbumartistId = `
CREATE INDEX IF NOT EXISTS index_album_artist ON albums(album_artist_id);
`
