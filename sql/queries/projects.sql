-- name: CreateProject :one
INSERT INTO projects (id, title, description, repo_url, live_url, image_url, is_featured, display_order)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, title, description, repo_url, live_url, image_url, is_featured, display_order, created_at, updated_at;

-- name: GetProjectByID :one
SELECT id, title, description, repo_url, live_url, image_url, is_featured, display_order, created_at, updated_at
FROM projects WHERE id = $1;

-- name: ListProjects :many
SELECT id, title, description, repo_url, live_url, image_url, is_featured, display_order, created_at, updated_at
FROM projects
WHERE
    ($1::boolean IS NULL OR is_featured = $1)
    AND ($2::text IS NULL OR title ILIKE '%' || $2::text || '%')
ORDER BY display_order ASC, created_at DESC
LIMIT $3 OFFSET $4;

-- name: CountProjects :one
SELECT COUNT(*) FROM projects
WHERE
    ($1::boolean IS NULL OR is_featured = $1)
    AND ($2::text IS NULL OR title ILIKE '%' || $2::text || '%');

-- name: UpdateProject :one
UPDATE projects
SET title = $2, description = $3, repo_url = $4, live_url = $5, image_url = $6,
    is_featured = $7, display_order = $8, updated_at = NOW()
WHERE id = $1
RETURNING id, title, description, repo_url, live_url, image_url, is_featured, display_order, created_at, updated_at;

-- name: DeleteProject :exec
DELETE FROM projects WHERE id = $1;

-- name: AddProjectSkill :exec
INSERT INTO project_skills (project_id, skill_id) VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveProjectSkill :exec
DELETE FROM project_skills WHERE project_id = $1 AND skill_id = $2;

-- name: RemoveAllProjectSkills :exec
DELETE FROM project_skills WHERE project_id = $1;

-- name: ListProjectSkills :many
SELECT s.id, s.name, s.category, s.proficiency, s.icon, s.display_order, s.created_at, s.updated_at
FROM skills s
JOIN project_skills ps ON ps.skill_id = s.id
WHERE ps.project_id = $1
ORDER BY s.display_order ASC, s.name ASC;
