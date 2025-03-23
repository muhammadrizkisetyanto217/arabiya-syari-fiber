CREATE TABLE IF NOT EXISTS user_evaluations (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    evaluation_id INTEGER NOT NULL,
    attempt INTEGER DEFAULT 1 NOT NULL,
    percentage_grade DOUBLE PRECISION DEFAULT 0 NOT NULL,
    time_duration INTEGER DEFAULT 0 NOT NULL,
    point INTEGER DEFAULT 0 NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_evaluation FOREIGN KEY(evaluation_id) REFERENCES evaluations(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_user_evaluations_user_id ON user_evaluations(user_id);
CREATE INDEX IF NOT EXISTS idx_user_evaluations_evaluation_id ON user_evaluations(evaluation_id);
