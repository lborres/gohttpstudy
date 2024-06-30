INSERT INTO MIGRATIONS (id, migration_name,start_datetime,end_datetime)
VALUES (
	'jesh3lfha2pymwebu53tk9vl','240427131639_cre_products',now(),now()
) ON CONFLICT(id) DO UPDATE SET
rolled_back_at = now()
;

CREATE SEQUENCE IF NOT EXISTS product_seq
AS BIGINT
INCREMENT BY 1
START 100
;

CREATE TYPE product_curr AS ENUM ('PHP', 'USD', 'JPY', 'CHF');

-- PK id, curr?
CREATE TABLE PRODUCTS (
	id                  BIGINT PRIMARY KEY DEFAULT nextval('product_seq'),
	name        		varchar(255) NOT NULL,
    description 		varchar(255) NOT NULL,
    sell_price        	NUMERIC(10,4) NOT NULL,
	currency            product_curr NOT NULL DEFAULT 'PHP',
	stock				INT NOT NULL DEFAULT 0,
	create_datetime	    timestamptz DEFAULT now(),
    update_datetime     timestamptz DEFAULT now(),

	UNIQUE (name)
);
