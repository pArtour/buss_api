 -- name: ListStopsByPlace :many
SELECT * FROM stop
WHERE to_tsvector('english', area) @@ to_tsquery('english', $1);