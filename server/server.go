package server

import (
	"github.com/gorilla/mux"
	"movie_api/controller"
	"movie_api/helper"
	"movie_api/repository"
	"movie_api/services"
	"net/http"
)

type Server struct {
	sql        repository.Repository
	services   services.Repository
	controller controller.Repository
}

func Register() *Server {
	SVR := &Server{}

	SVR.sql = repository.NewDbClient()

	SVR.services = services.NewServices(SVR.sql)

	SVR.controller = controller.NewController(SVR.services)

	return SVR
}

func (s *Server) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user", s.controller.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/api/users", s.controller.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/api/user", s.controller.GetUserById).Methods(http.MethodGet)

	err := http.ListenAndServe(":9000", r)
	helper.PrintError(err)
}
