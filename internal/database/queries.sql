-- name: ListTodos :many
SELECT id, description, done, created_at
FROM todos
WHERE done = false
ORDER BY id;
