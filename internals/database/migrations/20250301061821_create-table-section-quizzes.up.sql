CREATE TABLE user_quizzes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    section_quiz_id INT NOT NULL,
    attempt INT NOT NULL DEFAULT 1,
    percentage_grade FLOAT NOT NULL DEFAULT 0,
    time_duration INT NOT NULL DEFAULT 0,
    point INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_section_quiz FOREIGN KEY (section_quiz_id) REFERENCES section_quizzes(id) ON DELETE CASCADE,
    UNIQUE (user_id, section_quiz_id)
);
