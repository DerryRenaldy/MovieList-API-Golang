package repository

import (
	"context"
	"database/sql"
	"movie_api/models/request"
	"movie_api/models/response"
)

type dbClient struct {
	DB *sql.DB
}

type Repository interface {
	InsertUser(ctx context.Context, tx *sql.Tx, req request.InsertUsers) (*response.InsertUsers, error)
	GetAllUser(ctx context.Context) []response.Users
	GetUserById(ctx context.Context, id int)
}
