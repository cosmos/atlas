ALTER TABLE modules
ADD COLUMN IF NOT EXISTS stars INT DEFAULT 0;