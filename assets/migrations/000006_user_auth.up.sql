BEGIN;
CREATE TABLE user_auth (
    id UUID PRIMARY KEY, 
    user_id UUID NOT NULL REFERENCES public.user(user_id) ON DELETE CASCADE,     
    session_id UUID NOT NULL UNIQUE, -- Unique session identifier
    jwt_token_hash TEXT NOT NULL, 
    refresh_token_hash TEXT NOT NULL,
    device_details JSONB, -- Stores device-related data (e.g., browser, IP, OS)
    OTP INT NOT NULL,
    OTP_created_at TIMESTAMP NOT NULL,
    OTP_validity_secs INTEGER NOT NULL,
    OTP_attempts INT NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    -- Timestamp information
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMIT;