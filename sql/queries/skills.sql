-- name: CreateSkill :one
INSERT INTO skills (
    id,
    name,
    category,
    proficiency,
    icon,
    display_order
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, category, proficiency, icon, display_order, created_at, updated_at;

-- name: GetSkillByID :one
SELECT
    id,
    name,
    category,
    proficiency,
    icon,
    display_order,
    created_at,
    updated_at
FROM skills
WHERE id = $1;

-- name: GetSkillByName :one
SELECT 
    id,
    name,
    category,
    proficiency,
    icon,
    display_order,
    created_at,
    updated_at
FROM skills
WHERE name = $1;

-- name: ListSkills :many
SELECT
    id,
    name,
    category,
    proficiency,
    icon,
    display_order,
    created_at,
    updated_at
FROM skills
WHERE
    ($1::text IS NULL OR category = $1)
    AND ($2::text IS NULL OR name ILIKE '%' || $2::text || '%')
ORDER BY display_order ASC, name ASC
LIMIT $3 OFFSET $4;

-- name: CountSkills :one
SELECT COUNT(*)
FROM skills
WHERE
    ($1::text IS NULL OR category = $1)
    AND ($2::text IS NULL OR name ILIKE '%' || $2::text || '%');

-- name: UpdateSkill :one
UPDATE skills
SET
    name = $2,
    category = $3,
    proficiency = $4,
    icon = $5,
    display_order = $6,
    updated_at = NOW()
WHERE id = $1
RETURNING id, name, category, proficiency, icon, display_order, created_at, updated_at;

-- name: DeleteSkill :exec
DELETE FROM skills
WHERE id = $1;
