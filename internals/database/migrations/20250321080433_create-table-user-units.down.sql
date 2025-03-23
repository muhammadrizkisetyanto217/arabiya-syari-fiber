-- Drop indexes if exist
DROP INDEX IF EXISTS idx_user_units_user_id;
DROP INDEX IF EXISTS idx_user_units_unit_id;
DROP INDEX IF EXISTS idx_user_units_deleted_at;

-- Drop table
DROP TABLE IF EXISTS user_units;
