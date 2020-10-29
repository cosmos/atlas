BEGIN;
ALTER TABLE user_tokens
ADD COLUMN name VARCHAR NOT NULL;
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_tokens_user_id_name ON user_tokens(user_id, name);
COMMIT;
