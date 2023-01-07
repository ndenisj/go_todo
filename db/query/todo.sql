-- name: CreateTodo :one
INSERT INTO todos (
    owner, 
    title, 
    content
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1
LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTodo :one
UPDATE todos 
SET 
    title = CASE
        WHEN @set_title::boolean = TRUE THEN @title
        ELSE title
    END,
    content = CASE
        WHEN @set_content::boolean = TRUE THEN @content
        ELSE content
    END
WHERE id = @id
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos 
WHERE id = $1;