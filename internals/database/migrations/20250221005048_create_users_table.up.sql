CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    google_id VARCHAR(255) UNIQUE,
    role VARCHAR(20) NOT NULL DEFAULT 'user' CHECK (role IN ('owner', 'user', 'teacher', 'treasurer', 'admin')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

