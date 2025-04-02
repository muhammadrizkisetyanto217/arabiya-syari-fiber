-- +migrate Down
DROP INDEX IF EXISTS idx_user_quizzes_userid_quizid;
DROP INDEX IF EXISTS idx_user_quizzes_user_id;

DROP INDEX IF EXISTS idx_user_point_level_rank_user_id;
