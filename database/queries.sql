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
