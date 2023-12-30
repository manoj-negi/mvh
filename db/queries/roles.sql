-- name: CreateRole :one
INSERT INTO role (
    name,
    description,
    is_deleted
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetRole :one
SELECT * FROM role WHERE id = $1;

-- name: GetAllRoles :many
SELECT * FROM role;

-- name: UpdateRole :one
UPDATE role
SET
    name = $2,
    description = $3,
    is_deleted = $4
WHERE id = $1
RETURNING *;

-- name: DeleteRole :one
DELETE FROM role WHERE id = $1
RETURNING *;
