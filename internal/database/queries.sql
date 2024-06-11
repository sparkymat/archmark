-- name: CreateUser :one
INSERT INTO users (
  name, email, encrypted_password
) VALUES (
  @name::text, @email::text, @encrypted_password::text
) RETURNING *;

-- name: FetchUserByEmail :one
SELECT * FROM users
WHERE email = @email::text LIMIT 1;

-- name: FetchUserByID :one
SELECT * FROM users
WHERE id = @user_id::uuid LIMIT 1;

-- name: CreateBookmark :one
INSERT INTO bookmarks
(title, url, description, image, user_id)
VALUES
(@title::text, @url::text, @description::text, @image::text, @user_id::uuid)
RETURNING *;

-- name: SearchBookmarks :many
SELECT *
  FROM bookmarks t
  WHERE t.title ILIKE '%' || @query::text || '%'
    AND t.user_id = @user_id::uuid
  ORDER BY t.title ASC
  LIMIT @page_limit::int
  OFFSET @page_offset::int;

-- name: CountSearchedBookmarks :one
SELECT COUNT(id)
  FROM bookmarks t
  WHERE t.title ILIKE '%' || @query::text || '%'
    AND t.user_id = @user_id::uuid;

-- name: FetchRecentBookmarks :many
SELECT *
  FROM bookmarks t
  WHERE t.user_id = @user_id::uuid
  ORDER BY t.updated_at DESC
  LIMIT @page_limit::int
  OFFSET @page_offset::int;

-- name: CountRecentBookmarks :one
SELECT COUNT(id)
  FROM bookmarks t
  WHERE t.user_id = @user_id::uuid
;

-- name: FetchBookmarkByID :one
SELECT *
  FROM bookmarks t
  WHERE id = @id::uuid
  LIMIT 1;

-- name: FetchBookmarksByIDs :many
SELECT *
  FROM bookmarks t
  WHERE id = ANY(@ids::uuid[]);

-- name: DeleteBookmark :exec
DELETE FROM bookmarks t
  WHERE id = @id::uuid;

-- name: UpdateBookmarkTitle :one
UPDATE bookmarks t
SET title = @title::text
WHERE id = @id::uuid
RETURNING *;

-- name: UpdateBookmarkUrl :one
UPDATE bookmarks t
SET url = @url::text
WHERE id = @id::uuid
RETURNING *;

-- name: UpdateBookmarkDescription :one
UPDATE bookmarks t
SET description = @description::text
WHERE id = @id::uuid
RETURNING *;

-- name: UpdateBookmarkTranscript :one
UPDATE bookmarks t
SET transcript = @transcript::text
WHERE id = @id::uuid
RETURNING *;

-- name: UpdateBookmarkImage :one
UPDATE bookmarks t
SET image = sqlc.narg('image')
WHERE id = @id::uuid
RETURNING *;
