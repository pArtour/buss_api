-- name: GetTripById :one
SELECT * FROM trip
WHERE id = $1;