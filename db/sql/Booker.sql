-- name: CreateBookerAccount :one
INSERT INTO Booker (
  Name,
  Email,
  Phone_Number,
  Password,
  Address
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetBookerAccount :one
SELECT * FROM Booker
WHERE id = $1 LIMIT 1;

-- name: GetBookerAccountForUpdate :one
SELECT * FROM Booker
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListBookerAccount :many
SELECT * FROM Booker
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: UpdateBookerPassword :one
UPDATE Booker
SET Password = $4
WHERE id = $1
RETURNING *;

-- name: DeleteBookerAccount :exec
DELETE FROM Booker
WHERE id = $1;