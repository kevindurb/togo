-- name: ListTodos :many
SELECT id, description, done, created_at
FROM todos
WHERE done = false
AND user_id = ?
ORDER BY id;

-- name: CreateTodo :exec
INSERT INTO todos (description, user_id, done)
VALUES (?, ?, false)
