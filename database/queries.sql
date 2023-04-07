-- name: FetchUserByUsername :one
SELECT u.*
  FROM users u
  WHERE u.username = @email::text LIMIT 1;

-- name: FetchBookmarksList :many
SELECT b.*
  FROM bookmarks b
  WHERE b.user_id = @user_id::bigint
  ORDER BY b.created_at DESC
  LIMIT @page_limit::int
  OFFSET @page_offset::int;

-- name: CountBookmarksList :one
SELECT COUNT(*)
  FROM bookmarks b
  WHERE b.user_id = @user_id::bigint;

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

-- name: MarkBookmarkFetched :exec
UPDATE bookmarks
  SET status = 'fetched' AND file_path = @file_path::text
  WHERE id = @id::bigint;
