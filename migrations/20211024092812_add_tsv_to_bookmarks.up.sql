ALTER TABLE bookmarks
  ADD COLUMN weighted_tsv tsvector;

UPDATE bookmarks
  SET weighted_tsv = v.weighted_tsv
  FROM (
    SELECT id,
           setweight(to_tsvector('english', COALESCE(title,'')), 'A') ||
           setweight(to_tsvector('english', COALESCE(content,'')), 'B')
           AS weighted_tsv
     FROM bookmarks
  ) AS v
  WHERE v.id = bookmarks.id;

CREATE FUNCTION bookmarks_weighted_tsv_trigger() RETURNS trigger AS $$  
begin
  new.weighted_tsv :=
     setweight(to_tsvector('english', COALESCE(new.title,'')), 'A') ||
     setweight(to_tsvector('english', COALESCE(new.content,'')), 'B');
  return new;
end  
$$ LANGUAGE plpgsql;

CREATE TRIGGER bookmarks_upd_tsvector BEFORE INSERT OR UPDATE  
ON bookmarks
FOR EACH ROW EXECUTE PROCEDURE bookmarks_weighted_tsv_trigger();  

CREATE INDEX weighted_tsv_idx ON bookmarks USING GIST (weighted_tsv);
