CREATE SCHEMA IF NOT EXISTS todo_list;
CREATE TABLE IF NOT EXISTS todo_list.tasks (
	hash_key uuid DEFAULT gen_random_uuid() NOT NULL,
	"name" varchar NOT NULL,
	description varchar NULL,
	created timestamptz DEFAULT now() NOT NULL,
	updated timestamptz NULL,
	deadline timestamptz NOT NULL,
	closed bool DEFAULT false NOT NULL,
	CONSTRAINT tasks_pk PRIMARY KEY (hash_key),
	CONSTRAINT tasks_unique UNIQUE (name)
);
CREATE ROLE db_write PASSWORD 'write_pass' LOGIN;
CREATE ROLE db_read PASSWORD 'read_pass' LOGIN;
GRANT USAGE ON SCHEMA todo_list TO db_write;
GRANT USAGE ON SCHEMA todo_list TO db_read;
GRANT SELECT, INSERT, UPDATE, DELETE, TRUNCATE ON todo_list.tasks TO db_write;
GRANT SELECT ON todo_list.tasks TO db_read;