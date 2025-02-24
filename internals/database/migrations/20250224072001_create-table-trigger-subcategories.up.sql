CREATE OR REPLACE FUNCTION update_total_themes_or_levels()
RETURNS TRIGGER AS $$
BEGIN
    -- Jika DELETE, perbarui subcategory lama
    IF TG_OP = 'DELETE' THEN
        UPDATE subcategories
        SET total_themes_or_levels = (
            SELECT COUNT(*) FROM themes_or_levels WHERE subcategory_id = OLD.subcategory_id
        )
        WHERE id = OLD.subcategory_id;

    -- Jika UPDATE, cek apakah subcategory_id berubah
    ELSIF TG_OP = 'UPDATE' THEN
        -- Perbarui subcategory lama
        UPDATE subcategories
        SET total_themes_or_levels = (
            SELECT COUNT(*) FROM themes_or_levels WHERE subcategory_id = OLD.subcategory_id
        )
        WHERE id = OLD.subcategory_id;

        -- Jika subcategory_id berubah, perbarui juga subcategory baru
        IF OLD.subcategory_id IS DISTINCT FROM NEW.subcategory_id THEN
            UPDATE subcategories
            SET total_themes_or_levels = (
                SELECT COUNT(*) FROM themes_or_levels WHERE subcategory_id = NEW.subcategory_id
            )
            WHERE id = NEW.subcategory_id;
        END IF;

    -- Jika INSERT, perbarui subcategory baru
    ELSE
        UPDATE subcategories
        SET total_themes_or_levels = (
            SELECT COUNT(*) FROM themes_or_levels WHERE subcategory_id = NEW.subcategory_id
        )
        WHERE id = NEW.subcategory_id;
    END IF;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;
