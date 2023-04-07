CREATE TYPE bookmark_status AS ENUM ('pending', 'fetched', 'archived');
CREATE TABLE bookmarks (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    url TEXT NOT NULL,
    title TEXT,
    html TEXT,
    file_path TEXT,
    status bookmark_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
CREATE TRIGGER bookmarks_updated_at
  BEFORE UPDATE
  ON bookmarks
  FOR EACH ROW
    EXECUTE FUNCTION moddatetime(updated_at);
