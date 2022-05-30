 -- name: ListStopsByPlace :many
SELECT * FROM stop
WHERE area LIKE $1;