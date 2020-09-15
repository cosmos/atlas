BEGIN;
-- create modules_keywords many-to-many relationship table
CREATE TABLE modules_keywords (
  module_id int NOT NULL,
  keyword_id int NOT NULL,
  PRIMARY KEY (module_id, keyword_id),
  FOREIGN KEY (module_id) REFERENCES modules(id) ON UPDATE CASCADE,
  FOREIGN KEY (keyword_id) REFERENCES keywords(id) ON UPDATE CASCADE
);
CREATE INDEX IF NOT EXISTS keyword_id_idx ON modules_keywords(keyword_id);
COMMIT;
