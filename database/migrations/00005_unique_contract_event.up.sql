ALTER TABLE contract_events ADD CONSTRAINT unique_contract_event UNIQUE (contract_id, event_id);
