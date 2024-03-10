-- Rename the 'name' column to 'username'
ALTER TABLE users RENAME COLUMN name TO username;

-- Add a uniqueness constraint to 'username'
ALTER TABLE users ADD CONSTRAINT username_unique UNIQUE (username);