CREATE TABLE user_section_quizzes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    section_quizzes_id INTEGER NOT NULL,
    complete_quiz INTEGER[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
