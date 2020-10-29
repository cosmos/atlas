BEGIN;
ALTER TABLE user_tokens DROP COLUMN name;
DROP INDEX IF EXISTS idx_user_tokens_user_id_name;
COMMIT;
