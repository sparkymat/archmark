ALTER TABLE bookmarks ADD COLUMN category TEXT DEFAULT '';
CREATE INDEX bookmarks_category_idx ON bookmarks USING HASH (category);
