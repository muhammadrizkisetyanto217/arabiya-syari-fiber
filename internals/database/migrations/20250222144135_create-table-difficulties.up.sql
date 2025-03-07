CREATE TABLE IF NOT EXISTS difficulties (
-- 20250222144135

/*************  ✨ Smart Paste 📚  *************/
/******  fa015b7e-ae89-45aa-9f09-73ddfb448935  *******/
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description_short VARCHAR(200),
    description_long VARCHAR(3000),
    total_categories INT,
    status VARCHAR(10) CHECK (status IN ('active', 'pending', 'archived')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
)