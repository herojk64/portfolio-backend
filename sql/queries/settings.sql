-- name: CreateSettings :one
INSERT INTO settings (
    key,
    value
)
VALUES ($1,$2)
RETURNING key, value
;

-- name: UpdateSettings :one
UPDATE settings
SET
    value = $1
WHERE 
    key = $2
RETURNING key, value, created_at, updated_at;

-- name: GetSettingsByKey :one
SELECT 
    key, 
    value,
    created_at,
    updated_at
FROM settings
WHERE 
    key = $1
    ;

-- name: GetSettings :many
SELECT 
    key,
    value,
    created_at,
    updated_at
FROM settings
WHERE 
    ($1::text IS NULL OR key ILIKE '%' || $1::text || '%')
    AND ($2::TIMESTAMPTZ IS NULL OR created_at >= $2)
    AND ($3::TIMESTAMPTZ IS NULL OR updated_at >= $3)
LIMIT $4 OFFSET $5;

-- name: DeleteSettings :exec
DELETE FROM settings
WHERE key = $1;


