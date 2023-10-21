package db

import (
	"context"
)

const createBooker = `INSERT INTO booker (
	name_booker,
	email,
	phone_number,
	pwd,
	address_booker
  ) VALUES (
	$1, $2, $3, $4, $5
  )
  RETURNING *;
`
type createBookerParam struct{
	Name string `json:"name_booker"`
	Email string `json:"email"`
	PhoneNum string `json:"phone_number"`
	Password string `json:"pwd"`
	Address string `json:"address_booker"`
}

func (q *Queries) CreateBookerAccount(ctx context.Context, arg createBookerParam) (Booker, error) {
	row := q.db.QueryRow(ctx, createBooker, arg.Name)
	var i Booker
	err := row.Scan(&i.ID, &i.Name, &i.Email, &i.PhoneNum, &i.Password, &i.Address)
	return i, err
}

const deleteBookerAccount = `DELETE FROM booker
WHERE id = $1;
`

func (q *Queries) DeleteBookerAccount(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteBookerAccount, id)
	return err
}

const getBookerAccount = `SELECT * FROM booker
WHERE id = $1 LIMIT 1;
`

func (q *Queries) GetBookerAccount(ctx context.Context, id int64) (Booker, error) {
	row := q.db.QueryRow(ctx, getBookerAccount, id)
	var i Booker
	err := row.Scan(&i.ID, &i.Name, &i.Email, &i.PhoneNum, &i.Address, &i.CreatedAt)
	return i, err
}

const getBookerAccountForUpdate = `SELECT * FROM booker
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;
`

func (q *Queries) GetBookerAccountForUpdate(ctx context.Context, id int64) (Booker, error) {
	row := q.db.QueryRow(ctx, getBookerAccountForUpdate, id)
	var i Booker
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

const listBookerAccounts = `SELECT * FROM Booking
WHERE name_booker = $1
ORDER BY id
LIMIT $4
OFFSET $5;
`

type ListBookerAccountsParams struct {
	Name  string `json:"name_booker"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListBookerAccounts(ctx context.Context, arg ListBookerAccountsParams) ([]Booker, error) {
	rows, err := q.db.Query(ctx, listBookerAccounts, arg.Name, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Booker{}
	for rows.Next() {
		var i Booker
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

const updateBookerAccount = `UPDATE booker
SET pwd = $4
WHERE id = $1
RETURNING *;
`

type UpdateBookerAccountParams struct {
	ID      int64 `json:"id"`
	PassWord string `json:"pwd"`
}

func (q *Queries) UpdateBookerPassword(ctx context.Context, arg UpdateBookerAccountParams) (Booker, error) {
	row := q.db.QueryRow(ctx, updateBookerAccount, arg.ID, arg.PassWord)
	var i Booker
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