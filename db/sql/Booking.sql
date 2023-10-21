-- name: CreateBooking :one
INSERT INTO booking (
  client_id,
  booker_id,
  restaurant_id,
  date_and_time,
  status_booking
) VALUES (
  $1::serial, $2::serial, $3::serial, $4, $5
)
RETURNING *;

-- name: GetBooking :one
SELECT * FROM booking
WHERE id = $1 LIMIT 1;

-- name: GetBookingForUpdate :one
SELECT * FROM booking
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListBookingAccount :many
SELECT * FROM Booking
WHERE name_booker = $1
ORDER BY id
LIMIT $4
OFFSET $5;

-- name: UpdateBooking :one
UPDATE Booking
SET booker_id = $2, status_booking = $5
WHERE id = $1
RETURNING *;

-- name: CancelBooking :exec
DELETE FROM Booking
WHERE id = $1;