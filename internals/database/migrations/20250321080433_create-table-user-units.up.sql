-- Create table if not exists
CREATE TABLE IF NOT EXISTS user_units (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    unit_id INTEGER NOT NULL,
    is_reading BOOLEAN DEFAULT FALSE,
    is_evaluation BOOLEAN DEFAULT FALSE,
    is_complete_section_quizzes BOOLEAN DEFAULT FALSE,
    total_section_quizzes INTEGER DEFAULT 0,
    grade_exam DOUBLE PRECISION,
    grade_result DOUBLE PRECISION,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_unit FOREIGN KEY(unit_id) REFERENCES units(id) ON DELETE CASCADE
);

-- Create indexes only if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_user_units_user_id') THEN
        CREATE INDEX idx_user_units_user_id ON user_units(user_id);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_user_units_unit_id') THEN
        CREATE INDEX idx_user_units_unit_id ON user_units(unit_id);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_user_units_deleted_at') THEN
        CREATE INDEX idx_user_units_deleted_at ON user_units(deleted_at);
    END IF;
END
$$;
