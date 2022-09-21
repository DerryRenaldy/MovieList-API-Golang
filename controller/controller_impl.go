package controller

import (
	"encoding/json"
	"movie_api/helper"
	"movie_api/models/request"
	"net/http"
)

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	customerCreateRequest := request.InsertUsers{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&customerCreateRequest)
	helper.PrintError(err)

	customerResponse, err := c.Services.CreateUser(r.Context(), &customerCreateRequest)
	helper.PrintError(err)

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(customerResponse)
	helper.PrintError(err)
}

func (c *Controller) GetAllUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) GetUserById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) GetMovieList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) GetScreeningList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
