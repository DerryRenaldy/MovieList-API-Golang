package services

import (
	"context"
	"movie_api/models/request"
	"movie_api/models/response"
	"movie_api/repository"
)

type Services struct {
	Repository repository.Repository
}

type Repository interface {
	CreateUser(ctx context.Context, req *request.InsertUsers) (*response.InsertUsers, error)
	GetAllUser(ctx context.Context) ([]response.Users, error)
	GetUserById(ctx context.Context, id int) (*response.Users, error)
	GetMovieList(ctx context.Context) ([]response.Movie, error)
	GetScreeningList(ctx context.Context, id int) ([]response.ScreeningData, error)
}

func NewServices(repository repository.Repository) *Services {
	return &Services{
		Repository: repository,
	}
}
