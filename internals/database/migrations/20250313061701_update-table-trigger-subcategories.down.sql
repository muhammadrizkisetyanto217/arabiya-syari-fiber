-- Hapus trigger lama
DROP TRIGGER IF EXISTS themes_or_levels_after_insert_update_delete ON themes_or_levels;
DROP FUNCTION IF EXISTS update_total_themes_or_levels;