-- name: CreatePermission :one
INSERT INTO permission (
    name,
    permission,
    is_deleted
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetPermission :one
SELECT * FROM permission WHERE id = $1;

-- name: GetAllPermissions :many
SELECT * FROM permission;

-- name: UpdatePermission :one
UPDATE permission
SET
    name = $2,
    permission = $3,
    is_deleted = $4
WHERE id = $1
RETURNING *;

-- name: DeletePermission :one
DELETE FROM permission WHERE id = $1
RETURNING *;
