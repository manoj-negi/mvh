-- name: CreatePark :one
INSERT INTO park (
    brand_id,
    name,
    country_id,
    validated,
    is_deleted
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetAllPark :many
SELECT * FROM park;

-- name: GetParksByBrand :many
SELECT * FROM park 
WHERE brand_id = $1 
ORDER BY name;
