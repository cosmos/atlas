BEGIN;
-- 
-- Create the users table
-- 
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    github_user_id INT UNIQUE,
    github_access_token VARCHAR UNIQUE,
    email VARCHAR UNIQUE,
    url VARCHAR,
    avatar_url VARCHAR,
    gravatar_id VARCHAR,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);
CREATE INDEX idx_users_deleted_at ON users(deleted_at timestamptz_ops);
-- 
-- Create the keywords table
-- 
CREATE TABLE IF NOT EXISTS keywords (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);
CREATE INDEX idx_keywords_deleted_at ON keywords(deleted_at timestamptz_ops);
CREATE INDEX idx_keywords_tsvector ON keywords USING GIN(to_tsvector('english', name));
-- 
-- Create the modules table
-- 
CREATE TABLE IF NOT EXISTS modules (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR,
    homepage VARCHAR,
    documentation VARCHAR,
    repo VARCHAR NOT NULL,
    team VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_modules_name_team ON modules(name, team);
CREATE INDEX IF NOT EXISTS idx_modules_team ON modules(team);
CREATE INDEX idx_modules_deleted_at ON modules(deleted_at timestamptz_ops);
CREATE INDEX idx_modules_tsvector ON modules USING GIN(
    to_tsvector(
        'english',
        name || ' ' || team || ' ' || description
    )
);
-- 
-- Create the module_versions table
-- 
CREATE TABLE IF NOT EXISTS module_versions (
    id SERIAL PRIMARY KEY,
    version VARCHAR NOT NULL,
    sdk_compat VARCHAR,
    module_id INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (module_id) REFERENCES modules(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_module_versions_version_module_id ON module_versions(version, module_id);
CREATE INDEX IF NOT EXISTS idx_module_versions_version ON module_versions(module_id);
CREATE INDEX idx_module_versions_deleted_at ON module_versions(deleted_at timestamptz_ops);
-- 
-- Create the bug_trackers table
-- 
CREATE TABLE IF NOT EXISTS bug_trackers (
    id SERIAL PRIMARY KEY,
    url VARCHAR,
    contact VARCHAR,
    module_id INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (module_id) REFERENCES modules(id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS module_id_idx ON bug_trackers(module_id);
CREATE INDEX idx_bug_trackers_deleted_at ON bug_trackers(deleted_at timestamptz_ops);
-- 
-- Create a relationship between the keywords and modules tables
-- 
CREATE TABLE module_keywords (
    keyword_id INT NOT NULL,
    module_id INT NOT NULL,
    PRIMARY KEY (keyword_id, module_id)
);
CREATE INDEX IF NOT EXISTS module_id_idx ON module_keywords(module_id);
-- 
-- Create a relationship between the users and modules tables as authors
-- 
CREATE TABLE module_authors (
    user_id INT NOT NULL,
    module_id INT NOT NULL,
    PRIMARY KEY (user_id, module_id)
);
CREATE INDEX IF NOT EXISTS idx_module_authors_module_id ON module_authors(module_id);
--
-- Create a relationship between the modules and users tables as owners
--
CREATE TABLE module_owners (
    user_id INT NOT NULL,
    module_id INT NOT NULL,
    PRIMARY KEY (user_id, module_id)
);
CREATE INDEX IF NOT EXISTS idx_module_owners_module_id ON module_owners(module_id);
COMMIT;
