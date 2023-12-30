-- name: CreateBrand :one
INSERT INTO brand (
    name,
    logo,
    website,
    validated,
    is_deleted
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetAllBrand :many
SELECT * FROM brand;