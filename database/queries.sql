-- name: FetchUserByUsername :one
SELECT u.*
  FROM users u
  WHERE u.username = @email::text LIMIT 1;

-- name: FetchBookmarksList :many
SELECT b.*
  FROM bookmarks b
  WHERE b.user_id = @user_id::bigint
    AND b.deleted_at IS NULL
  ORDER BY b.created_at DESC
  LIMIT @page_limit::int
  OFFSET @page_offset::int;

-- name: CountBookmarksList :one
SELECT COUNT(*)
  FROM bookmarks b
  WHERE b.user_id = @user_id::bigint
    AND b.deleted_at IS NULL;

-- name: CreateBookmark :one
INSERT INTO bookmarks (
  user_id, url
) VALUES (
  @user_id::bigint, @url::text
) RETURNING *;

-- name: FetchBookmarkByID :one
SELECT b.*
  FROM bookmarks b
  WHERE b.id = @id::bigint
  LIMIT 1;

-- name: UpdateBookmarkDetails :exec
UPDATE bookmarks
  SET title = @title::text, html = @html::text
  WHERE id = @id::bigint;

-- name: MarkBookmarkFetched :exec
UPDATE bookmarks
  SET status = 'fetched', file_path = @file_path::text
  WHERE id = @id::bigint;

-- name: CreateUser :one
INSERT INTO users (
  name, username, encrypted_password
) VALUES (
  @name::text, @username::text, @encrypted_password::text
) RETURNING *;

-- name: SearchBookmarks :many
SELECT b.*
  FROM bookmarks b
  WHERE b.user_id = @user_id::bigint
    AND b.html_ts @@ to_tsquery('english', @query::text)
    AND b.deleted_at IS NULL
  LIMIT @page_limit::int
  OFFSET @page_offset::int;

-- name: CountBookmarksSearchResults :one
SELECT COUNT(*)
  FROM bookmarks b
  WHERE b.user_id = @user_id::bigint
    AND b.html_ts @@ to_tsquery('english', @query::text)
    AND b.deleted_at IS NULL;

-- name: FetchCategories :many
SELECT DISTINCT(b.category)
  FROM bookmarks b
  WHERE b.category != '' AND b.user_id = @user_id AND b.deleted_at IS NULL
  ORDER BY b.category ASC;

-- name: UpdateBookmarkCategory :exec
UPDATE bookmarks
  SET category = @category::text
  WHERE id = @id::bigint;

-- name: FetchArchivedBookmarks :many
SELECT b.*
  FROM bookmarks b
  WHERE b.deleted_at IS NOT NULL
  ORDER BY b.deleted_at ASC;

-- name: ArchiveBookmark :exec
UPDATE bookmarks b
  SET deleted_at = now()
  WHERE id = @id::bigint;

-- name: UnarchiveBookmark :exec
UPDATE bookmarks b
  SET deleted_at = NULL
  WHERE id = @id::bigint;

-- name: DeleteBookmarks :exec
DELETE FROM bookmarks b
  WHERE b.deleted_at IS NOT NULL AND b.deleted_at < now() - ('1 hour'::interval * @agehours::int);
