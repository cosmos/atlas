BEGIN;
-- move relevant columns from modules table to module_versions table
ALTER TABLE modules DROP COLUMN IF EXISTS documentation;
ALTER TABLE modules DROP COLUMN IF EXISTS repo;
ALTER TABLE module_versions
ADD COLUMN IF NOT EXISTS documentation VARCHAR;
ALTER TABLE module_versions
ADD COLUMN IF NOT EXISTS repo VARCHAR NOT NULL;
-- add published_by column to the module_versions table
ALTER TABLE module_versions
ADD COLUMN IF NOT EXISTS published_by INT NOT NULL;
COMMIT;
