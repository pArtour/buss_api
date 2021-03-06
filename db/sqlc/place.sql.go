// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: place.sql

package db

import (
	"context"
)

const listPlaces = `-- name: ListPlaces :many
SELECT id, name FROM place
ORDER BY id
`

func (q *Queries) ListPlaces(ctx context.Context) ([]Place, error) {
	rows, err := q.db.QueryContext(ctx, listPlaces)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Place{}
	for rows.Next() {
		var i Place
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
