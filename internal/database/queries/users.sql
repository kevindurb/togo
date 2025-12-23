-- name: CreateUser :exec
INSERT INTO users (username, password_hash)
VALUES (?, ?);

-- name: GetUser :one
SELECT * FROM users WHERE username = ?
