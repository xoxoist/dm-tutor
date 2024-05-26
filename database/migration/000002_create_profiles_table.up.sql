CREATE TABLE profiles
(
    id              SERIAL PRIMARY KEY,                             -- Auto-incremented primary key
    user_id         INTEGER     NOT NULL,                           -- Assuming user_id will link to users table but without any constraint
    first_name      VARCHAR(50) NOT NULL,                           -- First name of the user
    last_name       VARCHAR(50) NOT NULL,                           -- Last name of the user
    bio             TEXT,                                           -- A short bio of the user
    profile_picture VARCHAR(255),                                   -- URL to the profile picture
    deleted_at      TIMESTAMP NULL,                                 -- Timestamp for soft deletion
    created_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp for creation
    updated_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP  -- Timestamp for last update
);

-- Adding index for columns that might be frequently searched or filtered
CREATE INDEX idx_profiles_user_id ON profiles (user_id);
CREATE INDEX idx_profiles_last_name ON profiles (last_name);
