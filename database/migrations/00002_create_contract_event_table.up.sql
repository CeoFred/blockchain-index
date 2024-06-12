CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    signature VARCHAR(255) NOT NULL, -- keccak256 hash of the event signature
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
