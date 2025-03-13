CREATE OR REPLACE FUNCTION update_total_themes_or_levels()
RETURNS TRIGGER AS $$
BEGIN
    -- Jika DELETE, perbarui subcategory lama
    IF TG_OP = 'DELETE' THEN
        UPDATE subcategories
        SET total_themes_or_levels = (
            SELECT COUNT(*) FROM themes_or_levels WHERE subcategories_id = OLD.subcategories_id
        )
        WHERE id = OLD.subcategories_id;

    -- Jika UPDATE, cek apakah subcategories_id berubah
    ELSIF TG_OP = 'UPDATE' THEN
        -- Perbarui subcategory lama
        UPDATE subcategories
        SET total_themes_or_levels = (
            SELECT COUNT(*) FROM themes_or_levels WHERE subcategories_id = OLD.subcategories_id
        )
        WHERE id = OLD.subcategories_id;

        -- Jika subcategories_id berubah, perbarui juga subcategory baru
        IF OLD.subcategories_id IS DISTINCT FROM NEW.subcategories_id THEN
            UPDATE subcategories
            SET total_themes_or_levels = (
                SELECT COUNT(*) FROM themes_or_levels WHERE subcategories_id = NEW.subcategories_id
            )
            WHERE id = NEW.subcategories_id;
        END IF;

    -- Jika INSERT, perbarui subcategory baru
    ELSE
        UPDATE subcategories
        SET total_themes_or_levels = (
            SELECT COUNT(*) FROM themes_or_levels WHERE subcategories_id = NEW.subcategories_id
        )
        WHERE id = NEW.subcategories_id;
    END IF;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Buat ulang trigger untuk themes_or_levels
CREATE TRIGGER themes_or_levels_after_insert_update_delete
AFTER INSERT OR UPDATE OR DELETE ON themes_or_levels
FOR EACH ROW
EXECUTE FUNCTION update_total_themes_or_levels();