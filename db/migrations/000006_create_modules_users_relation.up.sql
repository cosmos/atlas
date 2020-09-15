BEGIN;
-- add user_id column to modules table that the FK references
ALTER TABLE modules
ADD COLUMN author INT;
-- create a FK relationship with modules table
ALTER TABLE modules
ADD CONSTRAINT fk_author FOREIGN KEY (author) REFERENCES users(id) ON DELETE
SET NULL;
-- create index on author FK
CREATE INDEX IF NOT EXISTS author_idx ON modules(author);
-- create a many-to-many relationship mapping modules and users
CREATE TABLE modules_users (
  module_id int NOT NULL,
  user_id int NOT NULL,
  PRIMARY KEY (module_id, user_id),
  FOREIGN KEY (module_id) REFERENCES modules(id) ON UPDATE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE
);
CREATE INDEX IF NOT EXISTS user_id_idx ON modules_users(user_id);
COMMIT;
