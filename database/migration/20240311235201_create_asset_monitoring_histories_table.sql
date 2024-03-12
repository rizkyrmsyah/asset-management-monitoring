-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE asset_monitoring_histories (
    id SERIAL PRIMARY KEY,
    asset_id INTEGER,
    user_id INTEGER,
    status VARCHAR,
    notes VARCHAR,
    created_at timestamptz,
    updated_at timestamptz
)

-- +migrate StatementEnd