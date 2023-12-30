-- name: CreateCostType :one
INSERT INTO cost_type (
    description,
    is_deleted
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetAllCostType :many
SELECT * FROM cost_type;