-- Create test database
CREATE DATABASE poolapp_test;

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE poolapp TO postgres;
GRANT ALL PRIVILEGES ON DATABASE poolapp_test TO postgres;