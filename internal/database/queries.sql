-- name: FetchUserByEmail :one
SELECT * FROM users
WHERE email = @email::text LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  name, email, encrypted_password
) VALUES (
  @name::text, @email::text, @encrypted_password::text
) RETURNING *;
