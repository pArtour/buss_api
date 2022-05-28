// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: trip.sql

package db

import (
	"context"
)

const getTripById = `-- name: GetTripById :one
SELECT id, route_id, headsign, long_name, direction_code, wheelchair_accessible FROM trip
WHERE id = $1
`

func (q *Queries) GetTripById(ctx context.Context, id int32) (Trip, error) {
	row := q.db.QueryRowContext(ctx, getTripById, id)
	var i Trip
	err := row.Scan(
		&i.ID,
		&i.RouteID,
		&i.Headsign,
		&i.LongName,
		&i.DirectionCode,
		&i.WheelchairAccessible,
	)
	return i, err
}
