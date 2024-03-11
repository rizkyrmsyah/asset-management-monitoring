-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE asset_mutations (
    id SERIAL PRIMARY KEY,
    asset_id INTEGER,
    type VARCHAR,
    notes VARCHAR,
    created_at timestamptz,
    updated_at timestamptz
)

-- +migrate StatementEnd