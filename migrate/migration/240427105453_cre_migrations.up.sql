-- Create a migrations table to track all migrations applied on the database
-- This is intended to be Automated in the future
-- Just manual for now
CREATE TABLE MIGRATIONS (
    id                varchar(36) PRIMARY KEY,
    -- checksum          varchar(64), --Not sure how to implement this yet
    migration_name    varchar(255),
    start_datetime    timestamptz,
    end_datetime      timestamptz,
    rolled_back_at    timestamptz,
    logs              text
);

-- Insert into Migrations table
INSERT INTO MIGRATIONS (id, migration_name,start_datetime,end_datetime)
VALUES (
	'zqailsrvwp4rh76hdqypllvo','240427105453_cre_migrations',now(),now()
) ON CONFLICT(id) DO UPDATE SET
rolled_back_at = now()
;