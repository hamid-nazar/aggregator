
-- +goose Up

CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY,       -- UUID as CHAR(36)
    name VARCHAR(255) NOT NULL,    -- TEXT is replaced with VARCHAR
    created_at TIMESTAMP NOT NULL, -- Use TIMESTAMP
    updated_at TIMESTAMP NOT NULL  -- Use TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;
