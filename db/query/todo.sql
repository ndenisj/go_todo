-- name: CreateTodo :one
INSERT INTO todos (
    user_id,
    owner, 
    title, 
    content
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1
LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTodo :one
UPDATE todos 
SET 
    title = COALESCE(sqlc.narg(title), title),
    content = COALESCE(sqlc.narg(content), content)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos 
WHERE id = $1;