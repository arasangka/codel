// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: logs.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createLog = `-- name: CreateLog :one
INSERT INTO logs (
  message, flow_id, type
)
VALUES (
  $1, $2, $3
)
RETURNING id, message, created_at, flow_id, type
`

type CreateLogParams struct {
	Message string
	FlowID  pgtype.Int8
	Type    string
}

func (q *Queries) CreateLog(ctx context.Context, arg CreateLogParams) (Log, error) {
	row := q.db.QueryRow(ctx, createLog, arg.Message, arg.FlowID, arg.Type)
	var i Log
	err := row.Scan(
		&i.ID,
		&i.Message,
		&i.CreatedAt,
		&i.FlowID,
		&i.Type,
	)
	return i, err
}

const getLogsByFlowId = `-- name: GetLogsByFlowId :many
SELECT id, message, created_at, flow_id, type
FROM logs
WHERE flow_id = $1
ORDER BY created_at ASC
`

func (q *Queries) GetLogsByFlowId(ctx context.Context, flowID pgtype.Int8) ([]Log, error) {
	rows, err := q.db.Query(ctx, getLogsByFlowId, flowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Log
	for rows.Next() {
		var i Log
		if err := rows.Scan(
			&i.ID,
			&i.Message,
			&i.CreatedAt,
			&i.FlowID,
			&i.Type,
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
