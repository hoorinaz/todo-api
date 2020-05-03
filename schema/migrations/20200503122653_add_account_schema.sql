-- migrate:up
CREATE SCHEMA IF NOT EXISTS account;

-- migrate:down
DROP SCHEMA IF EXISTS account;
