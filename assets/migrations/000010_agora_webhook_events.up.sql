Begin;
CREATE TABLE agora_webhook_events (
    id SERIAL PRIMARY KEY,
    event_type VARCHAR(100) NOT NULL, -- Type of event (e.g., "user_joined", "call_started")
    event_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp of the event
    uid VARCHAR(100), -- User ID from Agora (if applicable)
    channel_name VARCHAR(255), -- Channel name for the session
    session_id VARCHAR(255), -- Unique session identifier
    app_id VARCHAR(100) NOT NULL, -- Agora App ID
    payload JSONB NOT NULL, -- Raw event payload in JSON format
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
commit;