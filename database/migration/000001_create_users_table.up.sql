CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,                              -- Auto-incremented primary key
    email      VARCHAR(255) NOT NULL UNIQUE,                    -- Email column with unique constraint
    password   VARCHAR(255) NOT NULL,                           -- Password column
    username   VARCHAR(50)  NOT NULL UNIQUE,                    -- Username column with unique constraint
    deleted_at TIMESTAMP NULL,                                  -- Timestamp for soft deletion
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp for last update
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP  -- Timestamp for creation
);

-- Adding index for columns that might be frequently searched or filtered
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_username ON users (username);

-- Index for deleted_at to quickly filter non-deleted users
CREATE INDEX idx_users_deleted_at ON users (deleted_at);
