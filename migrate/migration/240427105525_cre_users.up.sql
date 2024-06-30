INSERT INTO MIGRATIONS (id, migration_name,start_datetime,end_datetime)
VALUES (
	'rylk1w27yhrqnzbyqi2774yj','240427105525_cre_users',now(),now()
) ON CONFLICT(id) DO UPDATE SET
rolled_back_at = now()
;

CREATE TYPE user_role AS ENUM ('ADMIN', 'USER');

CREATE TABLE USERS (
	id              varchar(36) PRIMARY KEY,
	username        varchar(255),
	email           varchar(255),
	password        varchar(255),
    role            user_role NOT NULL DEFAULT 'USER',
	create_datetime	timestamptz,
    update_datetime timestamptz
);
