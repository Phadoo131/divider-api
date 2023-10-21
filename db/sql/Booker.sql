-- name: CreateBookerAccount :one
INSERT INTO booker (
  name_booker,
  email,
  phone_number,
  pwd,
  address_booker
) VALUES (
  $1, $2::VARCHAR(255), $3, $4, $5
)
RETURNING *;

-- name: GetBookerAccount :one
SELECT * FROM booker
WHERE id = $1 LIMIT 1;

-- name: GetBookerAccountForUpdate :one
SELECT * FROM booker
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListBookerAccount :many
SELECT * FROM booker
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: UpdateBookerPassword :one
UPDATE booker
SET pwd = $4::VARCHAR(20)
WHERE id = $1
RETURNING *;

-- name: DeleteBookerAccount :exec
DELETE FROM booker
WHERE id = $1;