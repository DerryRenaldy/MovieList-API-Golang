package controller

import (
	"movie_api/services"
	"net/http"
)

type Controller struct {
	Services services.Repository
}

type Repository interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetAllUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetMovieList(w http.ResponseWriter, r *http.Request)
	GetScreeningList(w http.ResponseWriter, r *http.Request)
}

func NewController(services services.Repository) *Controller {
	return &Controller{
		Services: services,
	}
}
