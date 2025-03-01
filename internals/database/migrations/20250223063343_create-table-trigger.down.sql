-- Hapus trigger pada tabel subcategories
DROP TRIGGER IF EXISTS subcategories_after_insert_update_delete ON subcategories;

-- Hapus function update_total_subcategories
DROP FUNCTION IF EXISTS update_total_subcategories;

-- Hapus trigger pada tabel categories
DROP TRIGGER IF EXISTS categories_after_insert_update_delete ON categories;

-- Hapus function update_total_categories
DROP FUNCTION IF EXISTS update_total_categories;
