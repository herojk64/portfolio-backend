-- name: CreateExperience :one
INSERT INTO experience (id, company, role, description, start_date, end_date, location)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, company, role, description, start_date, end_date, location, created_at, updated_at;

-- name: GetExperienceByID :one
SELECT id, company, role, description, start_date, end_date, location, created_at, updated_at
FROM experience WHERE id = $1;

-- name: ListExperience :many
SELECT id, company, role, description, start_date, end_date, location, created_at, updated_at
FROM experience
WHERE
    ($1::text IS NULL OR company ILIKE '%' || $1::text || '%')
    AND ($2::text IS NULL OR role ILIKE '%' || $2::text || '%')
ORDER BY created_at DESC
LIMIT $3 OFFSET $4;

-- name: CountExperience :one
SELECT COUNT(*) FROM experience
WHERE
    ($1::text IS NULL OR company ILIKE '%' || $1::text || '%')
    AND ($2::text IS NULL OR role ILIKE '%' || $2::text || '%');

-- name: UpdateExperience :one
UPDATE experience
SET company = $2, role = $3, description = $4, start_date = $5, end_date = $6,
    location = $7, updated_at = NOW()
WHERE id = $1
RETURNING id, company, role, description, start_date, end_date, location, created_at, updated_at;

-- name: DeleteExperience :exec
DELETE FROM experience WHERE id = $1;

-- name: AddExperienceSkill :exec
INSERT INTO experience_skills (experience_id, skill_id) VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveExperienceSkill :exec
DELETE FROM experience_skills WHERE experience_id = $1 AND skill_id = $2;

-- name: RemoveAllExperienceSkills :exec
DELETE FROM experience_skills WHERE experience_id = $1;

-- name: ListExperienceSkills :many
SELECT s.id, s.name, s.category, s.proficiency, s.icon, s.display_order, s.created_at, s.updated_at
FROM skills s
JOIN experience_skills es ON es.skill_id = s.id
WHERE es.experience_id = $1
ORDER BY s.display_order ASC, s.name ASC;
