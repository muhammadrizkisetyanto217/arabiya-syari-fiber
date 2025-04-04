CREATE TABLE IF NOT EXISTS question_saved (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    source_type_id INTEGER NOT NULL, -- 1 = quiz, 2 = evaluation, 3 = exam
    question_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
