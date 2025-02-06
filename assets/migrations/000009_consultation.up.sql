BEGIN;
CREATE TABLE Consultation (
    Consultation_ID UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    User_ID UUID NOT NULL,
    Consultant_ID UUID NOT NULL,
    Session_ID UUID,
    Consultation_Time_Secs Int,
    Consultation_Type VARCHAR(50) NOT NULL,
    Consultation_State VARCHAR(50) NOT NULL,
    User_Wait_Time_Secs INT,
    Agora_Channel VARCHAR(100),
    Created_At TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    Updated_At TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

    -- Foreign Keys (if needed)
    CONSTRAINT fk_user FOREIGN KEY (User_ID) REFERENCES public.user(User_ID) ON DELETE CASCADE,
    CONSTRAINT fk_consultant FOREIGN KEY (Consultant_ID) REFERENCES Consultant(Consultant_ID) ON DELETE CASCADE
);

COMMIT;