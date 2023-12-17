-- name: create-table-projects

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
