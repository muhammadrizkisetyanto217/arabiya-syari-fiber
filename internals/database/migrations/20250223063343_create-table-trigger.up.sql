-- Function untuk memperbarui total_categories pada tabel difficulties
CREATE OR REPLACE FUNCTION update_total_categories()
RETURNS TRIGGER AS $$
BEGIN
    -- Jika DELETE, perbarui difficulty lama
    IF TG_OP = 'DELETE' THEN
        UPDATE difficulties
        SET total_categories = (
            SELECT COUNT(*) FROM categories WHERE difficulty_id = OLD.difficulty_id
        )
        WHERE id = OLD.difficulty_id;

    -- Jika UPDATE, cek apakah difficulty_id berubah
    ELSIF TG_OP = 'UPDATE' THEN
        -- Perbarui difficulty lama
        UPDATE difficulties
        SET total_categories = (
            SELECT COUNT(*) FROM categories WHERE difficulty_id = OLD.difficulty_id
        )
        WHERE id = OLD.difficulty_id;

        -- Jika difficulty_id berubah, perbarui juga difficulty baru
        IF OLD.difficulty_id IS DISTINCT FROM NEW.difficulty_id THEN
            UPDATE difficulties
            SET total_categories = (
                SELECT COUNT(*) FROM categories WHERE difficulty_id = NEW.difficulty_id
            )
            WHERE id = NEW.difficulty_id;
        END IF;

    -- Jika INSERT, perbarui difficulty baru
    ELSE
        UPDATE difficulties
        SET total_categories = (
            SELECT COUNT(*) FROM categories WHERE difficulty_id = NEW.difficulty_id
        )
        WHERE id = NEW.difficulty_id;
    END IF;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Buat trigger untuk menjalankan function di atas pada tabel categories
DROP TRIGGER IF EXISTS categories_after_insert_update_delete ON categories;


CREATE TRIGGER categories_after_insert_update_delete
AFTER INSERT OR UPDATE OR DELETE ON categories
FOR EACH ROW
EXECUTE FUNCTION update_total_categories();



-- Function untuk memperbarui total_subcategories pada tabel categories
CREATE OR REPLACE FUNCTION update_total_subcategories()
RETURNS TRIGGER AS $$
BEGIN
    -- Jika DELETE, perbarui kategori lama
    IF TG_OP = 'DELETE' THEN
        UPDATE categories
        SET total_subcategories = (
            SELECT COUNT(*) FROM subcategories WHERE categories_id = OLD.categories_id
        )
        WHERE id = OLD.categories_id;

    -- Jika UPDATE, cek apakah categories_id berubah
    ELSIF TG_OP = 'UPDATE' THEN
        -- Perbarui kategori lama
        UPDATE categories
        SET total_subcategories = (
            SELECT COUNT(*) FROM subcategories WHERE categories_id = OLD.categories_id
        )
        WHERE id = OLD.categories_id;

        -- Jika categories_id berubah, perbarui juga kategori baru
        IF OLD.categories_id IS DISTINCT FROM NEW.categories_id THEN
            UPDATE categories
            SET total_subcategories = (
                SELECT COUNT(*) FROM subcategories WHERE categories_id = NEW.categories_id
            )
            WHERE id = NEW.categories_id;
        END IF;

    -- Jika INSERT, perbarui kategori baru
    ELSE UPDATE categories
        SET total_subcategories = (
            SELECT COUNT(*) FROM subcategories WHERE categories_id = NEW.categories_id
        )
        WHERE id = NEW.categories_id;
    END IF;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;


-- Buat trigger untuk menjalankan function di atas pada tabel subcategories
DROP TRIGGER IF EXISTS subcategories_after_insert_update_delete ON subcategories;

CREATE TRIGGER subcategories_after_insert_update_delete
AFTER INSERT OR UPDATE OR DELETE ON subcategories
FOR EACH ROW
EXECUTE FUNCTION update_total_subcategories();
