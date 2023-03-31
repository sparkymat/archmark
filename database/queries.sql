-- name: FetchUserByUsername :one
SELECT u.*
  FROM users u
  WHERE u.username = @email::text LIMIT 1;
