-- Membuat fungsi untuk memperbarui total_themes_or_levels di subcategories
CREATE OR REPLACE FUNCTION update_total_themes_or_levels()
RETURNS TRIGGER AS $$
BEGIN
    -- Jika DELETE, perbarui subcategory lama
    IF TG_OP = 'DELETE' THEN
        UPDATE subcategories
        SET total_themes_or_levels = (
            SELECT COUNT(*) FROM themes_or_levels WHERE subategories_id = OLD.subategories_id
        )
        WHERE id = OLD.subategories_id;

    -- Jika UPDATE, cek apakah subategories_id berubah
    ELSIF TG_OP = 'UPDATE' THEN
        -- Perbarui subcategory lama
        UPDATE subcategories
        SET total_themes_or_levels = (
            SELECT COUNT(*) FROM themes_or_levels WHERE subategories_id = OLD.subategories_id
        )
        WHERE id = OLD.subategories_id;

        -- Jika subategories_id berubah, perbarui juga subcategory baru
        IF OLD.subategories_id IS DISTINCT FROM NEW.subategories_id THEN
            UPDATE subcategories
            SET total_themes_or_levels = (
                SELECT COUNT(*) FROM themes_or_levels WHERE subategories_id = NEW.subategories_id
            )
            WHERE id = NEW.subategories_id;
        END IF;

    -- Jika INSERT, perbarui subcategory baru
    ELSE
        UPDATE subcategories
        SET total_themes_or_levels = (
            SELECT COUNT(*) FROM themes_or_levels WHERE subategories_id = NEW.subategories_id
        )
        WHERE id = NEW.subategories_id;
    END IF;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Membuat trigger untuk themes_or_levels
CREATE TRIGGER themes_or_levels_after_insert_update_delete
AFTER INSERT OR UPDATE OR DELETE ON themes_or_levels
FOR EACH ROW
EXECUTE FUNCTION update_total_themes_or_levels();
