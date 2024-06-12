CREATE TABLE event_logs (
    id UUID PRIMARY KEY,
    contract_address VARCHAR(42) NOT NULL,
    event_name VARCHAR(255) NOT NULL,
    block_number BIGINT,
    transaction_hash VARCHAR(66),
    log_index INT,
    data JSONB,
    topics JSONB,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
