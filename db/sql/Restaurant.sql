-- name: CreateRestaurant :one
INSERT INTO restaurant(
  name_restaurant,
  location_restaurant,
  contact_Number
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetRestaurantsInfo :one
SELECT * FROM restaurant
WHERE id = $1 LIMIT 1;

-- name: GetRestaurantForUpdate :one
SELECT * FROM restaurant
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListRestaurants :many
SELECT * FROM restaurant
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: DeleteRestaurant :exec
DELETE FROM restaurant
WHERE id = $1;