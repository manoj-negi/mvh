-- name: CreatePeriod :one
INSERT INTO period (
    period,
    is_deleted
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetAllPeriod :many
SELECT * FROM period;