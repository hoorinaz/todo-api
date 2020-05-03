-- migrate:up
CREATE TABLE IF NOT EXISTS account.users
(
    id              SERIAL
        PRIMARY KEY,
    username        VARCHAR(100),
    email           VARCHAR(100),
    password        VARCHAR,
    created_at      TIMESTAMP DEFAULT  now(),
    updated_at      TIMESTAMP DEFAULT now()
);

-- migrate:down
DROP TABLE IF EXISTS account.users;
