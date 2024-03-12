-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE assets (
    id SERIAL PRIMARY KEY,
    location_id INTEGER NOT NULL,
    name VARCHAR,
    code VARCHAR UNIQUE NOT NULL,
    in_date DATE,
    source VARCHAR,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz
)

-- +migrate StatementEnd