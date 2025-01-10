BEGIN;
CREATE TABLE action (
    action_id SERIAL PRIMARY KEY,
    action_name VARCHAR(255) NOT NULL,
    action_code VARCHAR(255) NOT NULL,
    action_description TEXT,
    permission_id INT REFERENCES permission(permission_id),
    version INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
COMMIT;