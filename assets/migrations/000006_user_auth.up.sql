BEGIN;
CREATE TABLE user_auth (
    id UUID PRIMARY KEY, 
    user_id UUID NOT NULL REFERENCES public.user(user_id) ON DELETE CASCADE,     
    session_id UUID NOT NULL UNIQUE, -- Unique session identifier
    jwt_token_hash TEXT, 
    refresh_token_hash TEXT,
    device_details JSONB, -- Stores device-related data (e.g., browser, IP, OS)
    OTP INT,
    OTP_created_at TIMESTAMP,
    OTP_validity_secs INTEGER,
    OTP_attempts INT,
    phone_number VARCHAR(20) NOT NULL,
    -- Timestamp information
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT;