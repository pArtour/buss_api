-- name: GetAgencyById :one
SELECT * FROM agency
WHERE id = $1;