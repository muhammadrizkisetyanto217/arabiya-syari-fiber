-- Drop indexes
DROP INDEX IF EXISTS idx_user_evaluations_user_id;
DROP INDEX IF EXISTS idx_user_evaluations_evaluation_id;

-- Drop table
DROP TABLE IF EXISTS user_evaluations;
