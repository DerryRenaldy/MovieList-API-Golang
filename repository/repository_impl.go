package repository

import (
	"context"
	"database/sql"
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
	//TODO implement me
	panic("implement me")
}
func (d *dbClient) GetUserById(ctx context.Context, id int) (*response.Users, error) {
	//TODO implement me
	panic("implement me")
}
func (d *dbClient) GetMovieList(ctx context.Context) ([]response.Movie, error) {
	//TODO implement me
	panic("implement me")
}
func (d *dbClient) GetScreeningList(ctx context.Context, id int) ([]response.ScreeningData, error) {
	//TODO implement me
	panic("implement me")
}
