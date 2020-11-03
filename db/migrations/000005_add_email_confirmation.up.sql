BEGIN;
ALTER TABLE users
ADD COLUMN IF NOT EXISTS email_confirmed BOOLEAN DEFAULT FALSE;
CREATE TABLE user_email_confirmations (
  user_id INT UNIQUE NOT NULL,
  token uuid UNIQUE NOT NULL
);
COMMIT;
