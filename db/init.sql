
CREATE TABLE IF NOT EXISTS result_tab (
	result_id serial NOT NULL,
	"key" VARCHAR(64) NOT NULL,
	"value" VARCHAR(64) NOT NULL,
	CONSTRAINT result_tab_pk PRIMARY KEY (result_id),
	CONSTRAINT result_tab_unique_key UNIQUE ("key")
);