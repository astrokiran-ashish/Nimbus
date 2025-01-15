CREATE TABLE consultant (
    consultant_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Unique identifier for the astrologer
    user_id UUID NOT NULL REFERENCES public.user(user_id) ON DELETE CASCADE,  -- Reference to the user table
    expertise varchar(255),                                   -- Astrologer's expertise (e.g., Vedic, Tarot)
    state varchar(50) NOT NULL CHECK (state IN ('Pending', 'Document Submission', 'Review', 'Approved', 'Rejected')),
    version INTEGER DEFAULT 1,
    chat_channel varchar(255),
    call_channel varchar(255),
    live_channel varchar(255),
    video_call_channel varchar(255),
    created_at TIMESTAMP DEFAULT now(),                        -- Timestamp of astrologer profile creation
    updated_at TIMESTAMP DEFAULT now()                         -- Timestamp of the last update
);