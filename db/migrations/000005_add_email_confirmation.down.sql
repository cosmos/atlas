BEGIN;
ALTER TABLE users DROP COLUMN IF EXISTS email_confirmed;
DROP TABLE IF EXISTS user_email_confirmations CASCADE;
COMMIT;
