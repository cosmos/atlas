BEGIN;
-- remove new columns from the module_versions table
ALTER TABLE module_versions DROP COLUMN IF EXISTS published_by;
ALTER TABLE module_versions DROP COLUMN IF EXISTS documentation;
ALTER TABLE module_versions DROP COLUMN IF EXISTS repo;
-- add original columns back to the modules table
ALTER TABLE modules
ADD COLUMN IF NOT EXISTS documentation VARCHAR;
ALTER TABLE modules
ADD COLUMN IF NOT EXISTS repo VARCHAR;
COMMIT;
