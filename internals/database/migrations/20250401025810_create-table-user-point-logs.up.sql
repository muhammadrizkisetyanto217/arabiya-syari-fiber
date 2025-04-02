-- +migrate Up
CREATE TABLE IF NOT EXISTS user_point_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    points INTEGER NOT NULL,
    source_type VARCHAR(50),     -- misalnya: 'quiz', 'exam', dll
    source_id INTEGER,           -- ID dari sumber quiz/exam
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user FOREIGN KEY (user_id)
        REFERENCES users(id) ON DELETE CASCADE
);

-- Index gabungan untuk optimasi query total poin per sumber
CREATE INDEX IF NOT EXISTS idx_user_point_logs_user_source
ON user_point_logs(user_id, source_type, source_id);

-- Index tambahan untuk user_id (query umum / lookup)
CREATE INDEX IF NOT EXISTS idx_user_point_logs_user_id
ON user_point_logs(user_id);
