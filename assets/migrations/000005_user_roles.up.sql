CREATE TABLE user_roles (
    user_id UUID NOT NULL REFERENCES public.user(user_id) ON DELETE CASCADE,
    role_id INT NOT NULL REFERENCES role(role_id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT now(),
    UNIQUE (user_id, role_id)
);