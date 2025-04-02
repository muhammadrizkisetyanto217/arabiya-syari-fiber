-- +migrate Up
CREATE INDEX IF NOT EXISTS idx_user_quizzes_userid_quizid ON user_quizzes(user_id, quiz_id);
CREATE INDEX IF NOT EXISTS idx_user_quizzes_user_id ON user_quizzes(user_id);


CREATE INDEX IF NOT EXISTS idx_user_point_level_rank_user_id ON user_point_level_rank(user_id);

