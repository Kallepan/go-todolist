CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    --user_uuid UUID NOT NULL
);

-- CREATE TABLE users(
--     email VARCHAR (255) UNIQUE NOT NULL,
--     user_id UUID PRIMARY KEY
-- );


-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- ALTER TABLE users ALTER COLUMN user_id SET DEFAULT uuid_generate_v4();
-- ALTER TABLE tasks ADD CONSTRAINT fk_user_task FOREIGN KEY (user_uuid) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE;
