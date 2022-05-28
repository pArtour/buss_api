// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: stop.sql

package db

import (
	"context"
)

const listStopsByPlace = `-- name: ListStopsByPlace :many
SELECT id, code, name, lat, lon, zone_id, alias, area, "desc", lest_x, lest_y, zone_name, authority FROM stop
WHERE to_tsvector('english', area) @@ to_tsquery('english', $1)
`

func (q *Queries) ListStopsByPlace(ctx context.Context, toTsquery string) ([]Stop, error) {
	rows, err := q.db.QueryContext(ctx, listStopsByPlace, toTsquery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Stop
	for rows.Next() {
		var i Stop
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Name,
			&i.Lat,
			&i.Lon,
			&i.ZoneID,
			&i.Alias,
			&i.Area,
			&i.Desc,
			&i.LestX,
			&i.LestY,
			&i.ZoneName,
			&i.Authority,
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