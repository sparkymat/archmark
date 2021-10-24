DROP INDEX weighted_tsv_idx;

DROP TRIGGER IF EXISTS bookmarks_upd_tsvector ON bookmarks;

DROP FUNCTION IF EXISTS bookmarks_weighted_tsv_trigger;

ALTER TABLE bookmarks
  DROP COLUMN weighted_tsv;
