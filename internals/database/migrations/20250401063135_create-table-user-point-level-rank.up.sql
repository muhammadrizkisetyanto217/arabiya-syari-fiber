CREATE TABLE IF NOT EXISTS user_point_level_rank (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL UNIQUE,
    amount_total_quiz INTEGER DEFAULT 0,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
