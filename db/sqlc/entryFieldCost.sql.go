// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: entryFieldCost.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEntryFieldCost = `-- name: CreateEntryFieldCost :one
INSERT INTO entry_fields_cost (
    cost_id,
    entry_field_id,
    is_deleted
) VALUES (
    $1,
    $2,
    $3
) RETURNING id, cost_id, entry_field_id, is_deleted, created_at, updated_at
`

type CreateEntryFieldCostParams struct {
	CostID       int32       `json:"cost_id"`
	EntryFieldID int32       `json:"entry_field_id"`
	IsDeleted    pgtype.Bool `json:"is_deleted"`
}

func (q *Queries) CreateEntryFieldCost(ctx context.Context, arg CreateEntryFieldCostParams) (EntryFieldsCost, error) {
	row := q.db.QueryRow(ctx, createEntryFieldCost, arg.CostID, arg.EntryFieldID, arg.IsDeleted)
	var i EntryFieldsCost
	err := row.Scan(
		&i.ID,
		&i.CostID,
		&i.EntryFieldID,
		&i.IsDeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllEntryFieldCost = `-- name: GetAllEntryFieldCost :many
SELECT id, cost_id, entry_field_id, is_deleted, created_at, updated_at FROM entry_fields_cost
`

func (q *Queries) GetAllEntryFieldCost(ctx context.Context) ([]EntryFieldsCost, error) {
	rows, err := q.db.Query(ctx, getAllEntryFieldCost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EntryFieldsCost{}
	for rows.Next() {
		var i EntryFieldsCost
		if err := rows.Scan(
			&i.ID,
			&i.CostID,
			&i.EntryFieldID,
			&i.IsDeleted,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}