CREATE TABLE verifications (
    id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    phone_number TEXT NOT NULL,
    otp TEXT NOT NULL,
    exp TIMESTAMPTZ
);

CREATE INDEX ON verifications(id);