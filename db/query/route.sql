-- name: GetRouteById :one
SELECT * FROM route
WHERE id = $1;