-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz
)

-- +migrate StatementEnd