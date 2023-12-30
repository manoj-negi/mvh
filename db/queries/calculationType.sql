-- name: CreateCalculationType :one
INSERT INTO calculation_type (
    description,
    is_deleted
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetAllCalculationType :many
SELECT * FROM calculation_type;