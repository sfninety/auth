CREATE TABLE users (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    phone_number TEXT NOT NULL,
    verified boolean NOT NULL DEFAULT false,
    onboarded boolean NOT NULL DEFAULT false,
    device_identifier TEXT,
    password_hash TEXT,
    created TIMESTAMPTZ,
    updated TIMESTAMPTZ,
)

CREATE INDEX ON users (id, phone_number);