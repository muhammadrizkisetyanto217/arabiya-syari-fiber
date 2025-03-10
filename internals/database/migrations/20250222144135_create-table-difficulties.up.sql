CREATE TABLE IF NOT EXISTS difficulties (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description_short VARCHAR(200),
    description_long VARCHAR(3000),
    total_categories INT,
    status VARCHAR(10) CHECK (status IN ('active', 'pending', 'archived')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
)