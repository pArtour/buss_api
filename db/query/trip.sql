-- name: ListTripsById :many
SELECT * FROM trip
WHERE trip_id = $1;