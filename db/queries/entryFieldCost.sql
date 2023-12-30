-- name: CreateEntryFieldCost :one
INSERT INTO entry_fields_cost (
    cost_id,
    entry_field_id,
    is_deleted
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: GetAllEntryFieldCost :many
SELECT * FROM entry_fields_cost;