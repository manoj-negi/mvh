-- name: CreateStaticValues :one
INSERT INTO static_values (
    cost_input_id,
    description,
    value,
    is_deleted
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;

-- name: GetAllStaticValues :many
SELECT * FROM static_values;