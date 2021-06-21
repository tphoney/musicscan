-- name: create-table-users

CREATE TABLE IF NOT EXISTS users (
 user_id            SERIAL PRIMARY KEY
,user_email         VARCHAR(250)
,user_password      VARCHAR(250)
,user_token         VARCHAR(250)
,user_name          VARCHAR(250)
,user_company       VARCHAR(250)
,user_admin         BOOLEAN
,user_blocked       BOOLEAN
,user_created       INTEGER
,user_updated       INTEGER
,user_authed        INTEGER
,UNIQUE(user_token)
-- ,UNIQUE(lower(user_email))
);
