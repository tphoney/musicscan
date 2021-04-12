-- name: create-table-members

CREATE TABLE IF NOT EXISTS members (
 member_project_id INTEGER
,member_user_id    INTEGER
,member_role       INTEGER
,PRIMARY KEY(member_project_id, member_user_id)
);

-- name: create-index-members-project-id

CREATE INDEX IF NOT EXISTS index_members_project ON members(member_project_id)

-- name: create-index-members-user-id

CREATE INDEX IF NOT EXISTS index_members_user ON members(member_user_id)
