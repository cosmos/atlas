BEGIN;
-- create bugs table
CREATE TABLE IF NOT EXISTS bugs (
  id SERIAL PRIMARY KEY,
  url VARCHAR NOT NULL,
  contact VARCHAR NOT NULL
);
-- add bug_id column to modules table that the FK references
ALTER TABLE modules
ADD COLUMN bug_id INT;
-- create a FK relationship with bugs table
ALTER TABLE modules
ADD CONSTRAINT fk_bug FOREIGN KEY (bug_id) REFERENCES bugs(id) ON DELETE
SET NULL;
-- create index on bugs FK
CREATE INDEX IF NOT EXISTS bugs_modules_idx ON modules(bug_id);
COMMIT;
