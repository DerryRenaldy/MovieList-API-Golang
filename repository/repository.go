package repository

import (
	"context"
	"database/sql"
	"movie_api/database"
	"movie_api/helper"
	"movie_api/models/request"
	"movie_api/models/response"
)

type dbClient struct {
	DB *sql.DB
}

type Repository interface {
	InsertUser(ctx context.Context, req *request.InsertUsers) (*response.InsertUsers, error)
	GetAllUser(ctx context.Context) ([]response.Users, error)
	GetUserById(ctx context.Context, id int) (*response.Users, error)
	GetMovieList(ctx context.Context) ([]response.Movie, error)
	GetScreeningList(ctx context.Context, id int) ([]response.ScreeningData, error)
}

func NewDbClient() *dbClient {
	db, err := database.NewDB()
	helper.PrintError(err)

	return &dbClient{
		DB: db,
	}
}
