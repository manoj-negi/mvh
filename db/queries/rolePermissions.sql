-- name: CreateRolePermission :one
INSERT INTO roles_permission (
    role_id,
    permission_id,
    is_deleted
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetRolePermission :one
SELECT * FROM roles_permission WHERE id = $1;

-- name: GetAllRolePermissions :many
SELECT * FROM roles_permission;

-- name: UpdateRolePermission :one
UPDATE roles_permission
SET
    role_id = $2,
    permission_id = $3,
    is_deleted = $4
WHERE id = $1
RETURNING *;

-- name: DeleteRolePermission :one
DELETE FROM roles_permission WHERE id = $1
RETURNING *;
