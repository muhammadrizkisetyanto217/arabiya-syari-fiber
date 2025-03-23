-- Membuat tabel hanya jika belum ada
CREATE TABLE IF NOT EXISTS user_readings (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	reading_id INTEGER NOT NULL,
	is_reading BOOLEAN DEFAULT FALSE,
	created_at TIMESTAMPTZ DEFAULT NOW(),
	updated_at TIMESTAMPTZ DEFAULT NOW(),
	deleted_at TIMESTAMPTZ,

	-- Foreign key constraint
	CONSTRAINT fk_user_readings_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	CONSTRAINT fk_user_readings_reading FOREIGN KEY (reading_id) REFERENCES readings(id) ON DELETE CASCADE
);

-- Membuat index hanya jika belum ada
DO $$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_indexes WHERE indexname = 'idx_user_readings_deleted_at') THEN
		CREATE INDEX idx_user_readings_deleted_at ON user_readings(deleted_at);
	END IF;
END
$$;
