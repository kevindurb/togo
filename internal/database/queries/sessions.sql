-- name: CreateSession :exec
INSERT INTO sessions (id, user_id, expires_at)
VALUES (?, ?, ?);

-- name: GetSession :one
SELECT * FROM sessions WHERE id = ?
