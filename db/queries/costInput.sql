-- name: CreateCostInput :one
INSERT INTO cost_input (
    type,
    is_deleted
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetAllCostInput :many
SELECT * FROM cost_input;