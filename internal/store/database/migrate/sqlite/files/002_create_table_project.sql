-- name: create-table-projects

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
