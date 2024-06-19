CREATE TABLE event_logs (
    id UUID PRIMARY KEY,
    contract_id UUID NOT NULL,
    contract_event_id UUID NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    event_name VARCHAR(255) NOT NULL,
    block_number BIGINT,
    transaction_hash VARCHAR(66),
    log_index INT,
    data JSONB,
    topics JSONB,
    timestamp TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (transaction_hash, block_number)
);
