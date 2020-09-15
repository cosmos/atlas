CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    github_access_token VARCHAR NOT NULL,
    api_token VARCHAR NOT NULL
);
