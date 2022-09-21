package repository

import (
	"context"
	"database/sql"
	"errors"
	"movie_api/helper"
	"movie_api/models/request"
	"movie_api/models/response"
)

func (d *dbClient) InsertUser(ctx context.Context, req *request.InsertUsers) (*response.InsertUsers, error) {
	queryInsertUser := `INSERT INTO cinema_booking2.customers (first_name, last_name, password, email)
						VALUES (?,?,?,?);`

	query, err := d.DB.PrepareContext(ctx, queryInsertUser)
	helper.PrintError(err)

	defer func(query *sql.Stmt) {
		_ = query.Close()
	}(query)

	result, err := query.ExecContext(ctx, req.FirstName, req.LastName, req.Password, req.Email)
	helper.PrintError(err)

	id, err := result.LastInsertId()
	helper.PrintError(err)

	res := &response.InsertUsers{}
	responseQuery := "SELECT * FROM customers WHERE id = ?"
	err = d.DB.QueryRowContext(ctx, responseQuery, id).Scan(&res.ID, &res.FirstName, &res.LastName, &res.Password, &res.Email, &res.CreatedAt)
	helper.PrintError(err)

	return res, err

}

func (d *dbClient) GetAllUser(ctx context.Context) ([]response.Users, error) {
	queryAllUser := `SELECT first_name, last_name, password, email, created_at FROM customers;
`

	rows, err := d.DB.QueryContext(ctx, queryAllUser)
	helper.PrintError(err)

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var resSlice []response.Users

	for rows.Next() {
		res := response.Users{}

		err := rows.Scan(&res.FirstName, &res.LastName, &res.Password, &res.Email, &res.CreatedAt)
		helper.PrintError(err)

		resSlice = append(resSlice, res)

	}

	return resSlice, nil
}

func (d *dbClient) GetUserById(ctx context.Context, id int) (*response.Users, error) {
	queryGetById := `SELECT first_name, last_name, password, email, created_at FROM customers WHERE id = ?;`

	var res response.Users
	rows, err := d.DB.QueryContext(ctx, queryGetById, id)
	helper.PrintError(err)

	if rows.Next() {
		err = rows.Scan(&res.FirstName, &res.LastName, &res.Password, &res.Email, &res.CreatedAt)
		helper.PrintError(err)
		return &res, nil
	} else {
		return nil, errors.New("category is not found")
	}
}
func (d *dbClient) GetMovieList(ctx context.Context) ([]response.Movie, error) {
	//TODO implement me
	panic("implement me")
}
func (d *dbClient) GetScreeningList(ctx context.Context, id int) ([]response.ScreeningData, error) {
	//TODO implement me
	panic("implement me")
}
