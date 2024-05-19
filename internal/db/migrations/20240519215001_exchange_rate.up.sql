CREATE TABLE exchange_rate (
    id SERIAL PRIMARY KEY,
    exchange_rate NUMERIC,
    updated_at TIMESTAMP
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    last_email_sent_at TIMESTAMP
);