CREATE TABLE events (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    signature VARCHAR(255) NOT NULL, -- keccak256 hash of the event signature
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE events ADD CONSTRAINT unique_signature UNIQUE (signature);
