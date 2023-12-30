-- name: CreateEntryField :one
INSERT INTO entry_fields (
    calculation_type_id,
    user_email,
    calculation_date,
    park_id,
    asset_value,
    inventory_value,
    tax_asset_value,
    asset_increase_percentage,
    park_home_number,
    vat,
    ground,
    squarem2,
    constructions_year,
    renovation_percentage,
    yearly_revenue,
    yearly_revenue_increase_perc,
    asset_loan,
    loan_interest,
    is_deleted
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15,
    $16,
    $17,
    $18,
    $19
) RETURNING *;

-- name: GetAllEntryField :many
SELECT * FROM entry_fields;