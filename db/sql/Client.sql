-- name: CreateClientAccount :one
INSERT INTO Client (
  Name,
  Email,
  Phone_Number,
  Password,
  Address
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetClientAccount :one
SELECT * FROM Client
WHERE id = $1 LIMIT 1;

-- name: GetClientAccountForUpdate :one
SELECT * FROM Client
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListClientAccount :many
SELECT * FROM Client
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateClientPassword :one
UPDATE Client
SET Password = $4
WHERE id = $1
RETURNING *;

-- name: DeleteClientAccount :exec
DELETE FROM Client
WHERE id = $1;