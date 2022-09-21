package services

import (
	"context"
	"movie_api/helper"
	"movie_api/helper/password"
	"movie_api/models/request"
	"movie_api/models/response"
	"time"
)

func (s *Services) CreateUser(ctx context.Context, req *request.InsertUsers) (*response.InsertUsers, error) {
	// Hashing Password
	pass, err := password.HashPassword(req.Password)
	helper.PrintError(err)
	req.Password = pass

	res, err := s.Repository.InsertUser(ctx, req)
	helper.PrintError(err)

	res.CreatedAt = time.Now()
	res.Password = pass
	req.Password = pass

	return res, nil
}

func (s *Services) GetAllUser(ctx context.Context) ([]response.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Services) GetUserById(ctx context.Context, id int) (*response.Users, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Services) GetMovieList(ctx context.Context) ([]response.Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Services) GetScreeningList(ctx context.Context, id int) ([]response.ScreeningData, error) {
	//TODO implement me
	panic("implement me")
}
