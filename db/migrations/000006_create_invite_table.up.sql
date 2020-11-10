BEGIN;
CREATE TABLE module_owner_invites (
  module_id INT NOT NULL,
  invited_user_id INT NOT NULL,
  invited_by_user_id INT NOT NULL,
  token uuid UNIQUE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_module_owner_invites ON module_owner_invites(module_id, invited_user_id);
CREATE INDEX IF NOT EXISTS idx_module_owner_invites_invited_user_id ON module_owner_invites(invited_user_id);
CREATE INDEX IF NOT EXISTS idx_module_owner_invites_invited_by_user_id ON module_owner_invites(invited_by_user_id);
COMMIT;
