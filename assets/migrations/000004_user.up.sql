CREATE TABLE public.user (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Unique identifier (UUID)
    name TEXT,
    email TEXT UNIQUE,                          -- Unique email address
    dob DATE,                                            -- Date of birth
    phone_number TEXT UNIQUE,                            -- Phone number
    version INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),                  -- Timestamp when the user was created
    updated_at TIMESTAMP DEFAULT now()                   -- Timestamp when the user was last updated
);