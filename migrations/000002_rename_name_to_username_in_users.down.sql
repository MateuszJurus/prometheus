-- Remove the uniqueness constraint from 'username'
ALTER TABLE users DROP CONSTRAINT username_unique;

-- Rename the 'username' column back to 'name'
ALTER TABLE users RENAME COLUMN username TO name;