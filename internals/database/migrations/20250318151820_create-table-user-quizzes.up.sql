CREATE TABLE user_quizzes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    quiz_id INT NOT NULL, -- Ubah relasi ke quizzes
    attempt INT NOT NULL DEFAULT 1,
    percentage_grade FLOAT NOT NULL DEFAULT 0,
    time_duration INT NOT NULL DEFAULT 0,
    point INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_quiz FOREIGN KEY (quiz_id) REFERENCES quizzes(id) ON DELETE CASCADE, -- Relasi ke quizzes
    UNIQUE (user_id, quiz_id)
);


