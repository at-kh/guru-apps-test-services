-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION trigger_set_updated_at() RETURNS trigger AS $$
BEGIN
    NEW.updated_at = now();
RETURN NEW;
END
$$ LANGUAGE plpgsql;
-- +migrate StatementEnd

-- +migrate Down
DROP FUNCTION IF EXISTS trigger_set_updated_at CASCADE;
DROP EXTENSION IF EXISTS "pgcrypto";
