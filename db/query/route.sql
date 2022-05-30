-- name: GetRouteById :one
SELECT * FROM route
WHERE route_id = $1;