ALTER TABLE user_actions 
ADD COLUMN transaction_hash TEXT;

ALTER TABLE user_actions ADD CONSTRAINT unique_user_action UNIQUE (user_id,transaction_hash);