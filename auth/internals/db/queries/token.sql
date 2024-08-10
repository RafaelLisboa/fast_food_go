-- name: CreateRefreshToken :exec
INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3);


-- name: GetRefreshToken :one
SELECT token FROM refresh_tokens WHERE token = $1;