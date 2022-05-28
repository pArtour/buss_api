-- name: ListTimesByStop :many
SELECT * FROM stop_time
WHERE stop_id = $1;