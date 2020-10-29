BEGIN;
CREATE TABLE user_module_favorites (
  user_id INT NOT NULL,
  module_id INT NOT NULL,
  PRIMARY KEY (user_id, module_id)
);
CREATE INDEX IF NOT EXISTS idx_user_module_favorites_module_id ON user_module_favorites(module_id);
COMMIT;
