-- +goose Up
CREATE SCHEMA users;

SET
SEARCH_PATH TO users, PUBLIC;

CREATE TABLE users (
  id         text        NOT NULL,
  name       text        NOT NULL,
  email text        NOT NULL,
  enabled    bool        NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);



CREATE TABLE inbox (
  id          text        NOT NULL,
  name        text        NOT NULL,
  subject     text        NOT NULL,
  data        bytea       NOT NULL,
  metadata    bytea       NOT NULL,
  sent_at     timestamptz NOT NULL,
  received_at timestamptz NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE outbox (
  id           text        NOT NULL,
  name         text        NOT NULL,
  subject      text        NOT NULL,
  data         bytea       NOT NULL,
  metadata     bytea       NOT NULL,
  sent_at      timestamptz NOT NULL,
  published_at timestamptz,
  PRIMARY KEY (id)
);

CREATE INDEX users_unpublished_idx ON outbox (published_at) WHERE published_at IS NULL;

-- +goose Down
DROP SCHEMA IF EXISTS users CASCADE;
-- SET SEARCH_PATH TO users;
--
-- DROP TABLE IF EXISTS outbox;
-- DROP TABLE IF EXISTS inbox;
-- DROP TABLE IF EXISTS users;
