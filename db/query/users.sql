-- name: CreateUser :one
INSERT INTO users (
    full_name,
    phone, 
    username, 
    hashed_password,
    email
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users 
SET 
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    phone = COALESCE(sqlc.narg(phone), phone),
    username = COALESCE(sqlc.narg(username), username),
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password)
WHERE id = sqlc.arg(id)
RETURNING *;