--
-- USERS
--

INSERT INTO users VALUES
(1, 'jane@example.com', '', '12345', true, false, 1286668800, 1602374400, 1602460800),
(2, 'john@example.com', '', '54321', false, true, 1286668800, 1602374400, 1602460800);

--
-- PROJECTS
--

INSERT INTO projects VALUES (
 1
,'sourcegraph'
,'Sourcegraph makes code search universal so developers can work on solving problems.'
,'c81e728d9d4c2f636f067f89cc14862c'
,false
,1286668800
,1602374400
), (
 2
,'gitlab'
,'GitLab is a web-based open source Git repository manager with wiki and issue tracking features and built-in CI/CD.'
,'a87ff679a2f3e71d9181a67b7542122c'
,false
,1286668800
,1602374400
);

--
-- MEMBERS
--

INSERT INTO members VALUES
(1, 1, 0),
(1, 2, 1),
(2, 1, 1);
