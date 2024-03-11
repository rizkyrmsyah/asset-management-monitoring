-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE assets (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    code VARCHAR UNIQUE,
    in_date DATE,
    source VARCHAR,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz
)

-- +migrate StatementEnd