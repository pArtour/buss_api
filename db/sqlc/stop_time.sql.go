// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: stop_time.sql

package db

import (
	"context"
)

const listTimesByStop = `-- name: ListTimesByStop :many
SELECT trip_id, arrival_time, departure_time, stop_id, stop_sequence, pickup_type, drop_off_type FROM stop_time
WHERE stop_id = $1
`

func (q *Queries) ListTimesByStop(ctx context.Context, stopID int32) ([]StopTime, error) {
	rows, err := q.db.QueryContext(ctx, listTimesByStop, stopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []StopTime{}
	for rows.Next() {
		var i StopTime
		if err := rows.Scan(
			&i.TripID,
			&i.ArrivalTime,
			&i.DepartureTime,
			&i.StopID,
			&i.StopSequence,
			&i.PickupType,
			&i.DropOffType,
		); err != nil {
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
