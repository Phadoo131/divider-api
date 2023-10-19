-- name: CreateBooking :one
INSERT INTO Booking (
  ClientID,
  BookerID,
  RestaurantID,
  Date_and_Time,
  Status
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetBooking :one
SELECT * FROM Booking
WHERE id = $1 LIMIT 1;

-- name: GetBookingForUpdate :one
SELECT * FROM Booking
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListBookingAccount :many
SELECT * FROM Booking
ORDER BY id
LIMIT $3
OFFSET $5;

-- name: UpdateBooking :one
UPDATE Booking
SET BookerID = $2 AND SET Status = $5
WHERE id = $1
RETURNING *;

-- name: CancelBooking :exec
DELETE FROM Booking
WHERE id = $1;