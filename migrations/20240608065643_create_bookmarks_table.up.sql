CREATE EXTENSION IF NOT EXISTS moddatetime;

CREATE TABLE bookmarks (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  title text NOT NULL,
  url text NOT NULL,
  description text NOT NULL,
  transcript text NOT NULL DEFAULT '',
  image text,
  user_id uuid REFERENCES users(id) NOT NULL,
  created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TRIGGER bookmarks_updated_at
  BEFORE UPDATE
  ON bookmarks
  FOR EACH ROW
    EXECUTE FUNCTION moddatetime(updated_at);
