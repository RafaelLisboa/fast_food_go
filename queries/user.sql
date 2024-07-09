-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, username) 
VALUES ($1, $2, $3, $4, $5);

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;


-- name: GetUserByID :one
SELECT * FROM users WHERE username = $1;