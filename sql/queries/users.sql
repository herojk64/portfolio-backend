-- name: CreateUser :one
INSERT INTO users(
    id,
    name,
    username,
    email,
    avatar,
    password_hash,
    is_super
)
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING id, name, username, email, avatar, is_super, created_at;

-- name: GetUserByID :one
SELECT
    id,
    name,
    username,
    email,
    avatar,
    is_super,
    created_at,
    updated_at
FROM
    users
WHERE
    id = $1;

-- name: GetUserByUsername :one
SELECT
    id,
    name,
    username,
    email,
    avatar,
    is_super,
    created_at,
    updated_at
FROM
    users
WHERE
    username = $1;

-- name: UpdateUser :one
UPDATE users
SET 
    name = $2,
    username = $3,
    email = $4,
    avatar = $5,
    updated_at = NOW()
WHERE
    id = $1 
RETURNING id, name, username, email, avatar, is_super, updated_at;

-- name: UpdatePassword :one
UPDATE users
SET password_hash = $2
WHERE
    id = $1
RETURNING id, name, username, email, avatar, is_super, updated_at;

-- name: GiveSuperAccess :one
UPDATE users
SET is_super = $2
WHERE
    id = $1
RETURNING name, username, email, avatar, is_super, updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE
    id = $1;

-- name: SelectUsers :many
SELECT
    id,
    name,
    username,
    email,
    avatar,
    created_at,
    updated_at
FROM users
WHERE
    ($1::text IS NULL OR id LIKE $1::text || '%')
    AND ($2::text IS NULL OR name ILIKE '%' || $2::text || '%')
    AND ($3::text IS NULL OR username ILIKE '%' || $3::text || '%')
    AND ($4::text IS NULL OR email ILIKE '%' || $4::text || '%')
    AND ($5::boolean IS NULL OR is_super = $5)
    AND ($6::timestamptz IS NULL OR created_at >= $6)
    AND ($7::timestamptz IS NULL OR updated_at >= $7)
LIMIT $8 OFFSET $9;
