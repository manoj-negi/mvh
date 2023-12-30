-- name: CreateCostCalculation :one
INSERT INTO cost_calculation (
    parkcost,
    assetcost,
    loancost,
    tax_cost,
    asset_value_cost,
    revenue,
    total_cost,
    result,
    result_perc,
    saving_revenue_amount,
    saving_revenue_perc,
    yield_diff,
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
    $13
) RETURNING *;

-- name: GetAllCostCalculation :many
SELECT * FROM cost_calculation;