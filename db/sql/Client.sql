-- name: CreateClientAccount :one
INSERT INTO client (
  name_client,
  email,
  phone_number,
  pwd,
  address_client
) VALUES (
  $1, $2::VARCHAR(255), $3, $4, $5
)
RETURNING *;

-- name: GetClientAccount :one
SELECT * FROM client
WHERE id = $1 LIMIT 1;

-- name: GetClientAccountForUpdate :one
SELECT * FROM client
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListClientAccount :many
SELECT * FROM client
ORDER BY id;

-- name: UpdateClientPassword :one
UPDATE client
SET pwd = $4::VARCHAR(20)
WHERE id = $1
RETURNING *;

-- name: DeleteClientAccount :exec
DELETE FROM client
WHERE id = $1;