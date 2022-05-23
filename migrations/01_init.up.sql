CREATE TABLE IF NOT EXISTS users  (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    phone_number TEXT NOT NULL,
    verified boolean NOT NULL DEFAULT false,
    onboarded boolean NOT NULL DEFAULT false,
    device_identifier TEXT,
    password_hash TEXT,
    created TIMESTAMPTZ,
    updated TIMESTAMPTZ
);


CREATE TABLE IF NOT EXISTS verifications (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    phone_number TEXT NOT NULL,
    otp TEXT NOT NULL,
    expiry TIMESTAMPTZ
);