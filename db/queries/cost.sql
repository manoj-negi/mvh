-- name: CreateCost :one
INSERT INTO cost (
    cost_input_id,
    cost_type_id,
    period_id,
    title,
    is_deleted
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetAllCost :many
SELECT * FROM cost;