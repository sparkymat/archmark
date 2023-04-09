ALTER TABLE bookmarks ADD COLUMN html_ts tsvector
    GENERATED ALWAYS AS (to_tsvector('english', html)) STORED;
CREATE INDEX html_ts_idx ON bookmarks USING GIN (html_ts);
