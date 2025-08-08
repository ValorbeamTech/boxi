-- Database initialization script
-- This runs when the PostgreSQL container starts for the first time

-- Create databases
CREATE DATABASE myapp;
CREATE DATABASE myapp_dev;
CREATE DATABASE myapp_test;

-- Create user (optional, for better security)
DO
$do$
BEGIN
   IF NOT EXISTS (
      SELECT FROM pg_catalog.pg_roles 
      WHERE rolname = 'myapp_user') THEN
      
      CREATE ROLE myapp_user LOGIN PASSWORD 'myapp_password';
   END IF;
END
$do$;

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE myapp TO myapp_user;
GRANT ALL PRIVILEGES ON DATABASE myapp_dev TO myapp_user;
GRANT ALL PRIVILEGES ON DATABASE myapp_test TO myapp_user;

-- Grant privileges on schemas
\c myapp;
GRANT ALL ON SCHEMA public TO myapp_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO myapp_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO myapp_user;

\c myapp_dev;
GRANT ALL ON SCHEMA public TO myapp_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO myapp_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO myapp_user;

\c myapp_test;
GRANT ALL ON SCHEMA public TO myapp_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO myapp_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO myapp_user;