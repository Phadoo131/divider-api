-- name: CreateRestaurant :one
INSERT INTO Restaurant(
  Name,
  Location,
  Contact_Number
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetRestaurantsInfo :one
SELECT * FROM Restaurant
WHERE id = $1 LIMIT 1;

-- name: GetRestaurantForUpdate :one
SELECT * FROM Restaurant
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListRestaurants :many
SELECT * FROM Restaurant
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: DeleteRestaurant :exec
DELETE FROM Restaurant
WHERE id = $1;