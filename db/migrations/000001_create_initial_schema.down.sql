BEGIN;
-- 
-- Drop the users table
-- 
DROP TABLE IF EXISTS users CASCADE;
-- 
-- Drop the keywords table
-- 
DROP TABLE IF EXISTS keywords CASCADE;
-- 
-- Drop the modules table
-- 
DROP TABLE IF EXISTS modules CASCADE;
-- 
-- Drop the module_versions table
-- 
DROP TABLE IF EXISTS module_versions CASCADE;
-- 
-- Drop the bug_trackers table
-- 
DROP TABLE IF EXISTS bug_trackers CASCADE;
-- 
-- Drop the module_keywords table
-- 
DROP TABLE IF EXISTS module_keywords CASCADE;
-- 
-- Drop the module_authors table
-- 
DROP TABLE IF EXISTS module_authors CASCADE;
--
-- Drop the module_owners table
--
DROP TABLE IF EXISTS module_owners CASCADE;
--
-- Drop the user_tokens table
--
DROP TABLE IF EXISTS user_tokens CASCADE;
COMMIT;
