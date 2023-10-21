package db

import (
	"context"
)

const createClient = `INSERT INTO client (
	name,
	email,
	phone_number,
	password,
	address
  ) VALUES (
	$1, $2, $3, $4, $5
  )
  RETURNING *;
`
type createClientParam struct{
	Name string `json:"name_client"`
	Email string `json:"email"`
	PhoneNum string `json:"phone_number"`
	Password string `json:"pwd"`
	Address string `json:"address_client"`
}

func (q *Queries) CreateClientAccount(ctx context.Context, arg createClientParam) (Client, error) {
	row := q.db.QueryRow(ctx, createClient, arg.Name, arg.Email, arg.PhoneNum, arg.Password, arg.Address)
	var i Client
	err := row.Scan(&i.ID, &i.Name, &i.Email, &i.PhoneNum, &i.Password, &i.Address)
	return i, err
}

const deleteClientAccount = `DELETE FROM client
WHERE id = $1;
`

func (q *Queries) DeleteClientAccount(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteClientAccount, id)
	return err
}

const getClientAccount = `SELECT * FROM client
WHERE id = $1 LIMIT 1;
`

func (q *Queries) GetClientAccount(ctx context.Context, id int64) (Client, error) {
	row := q.db.QueryRow(ctx, getClientAccount, id)
	var i Client
	err := row.Scan(&i.ID, &i.Name, &i.Email, &i.PhoneNum, &i.Address, &i.CreatedAt)
	return i, err
}

const getClientrAccountForUpdate = `SELECT * FROM client
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;
`

func (q *Queries) GetClientAccountForUpdate(ctx context.Context, id int64) (Client, error) {
	row := q.db.QueryRow(ctx, getBookerAccountForUpdate, id)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PhoneNum,
		&i.Password,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const listClientAccounts = `SELECT * FROM client
WHERE name_client = $1
ORDER BY id
LIMIT $4
OFFSET $5;
`

type ListClientAccountsParams struct {
	Name  string `json:"name_client"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListClientAccounts(ctx context.Context, arg ListClientAccountsParams) ([]Client, error) {
	rows, err := q.db.Query(ctx, listClientAccounts, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Client{}
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.PhoneNum,
			&i.Password,
			&i.Address,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateClientAccount = `UPDATE client
SET pwd = $4
WHERE id = $1
RETURNING *;
`

type UpdateClientAccountParams struct {
	ID      int64 `json:"id"`
	PassWord string `json:"pwd"`
}

func (q *Queries) UpdateClientPassword(ctx context.Context, arg UpdateClientAccountParams) (Client, error) {
	row := q.db.QueryRow(ctx, updateClientAccount, arg.ID, arg.PassWord)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.PhoneNum,
		&i.Password,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}