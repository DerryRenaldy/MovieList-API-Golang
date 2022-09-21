package controller

import (
	"encoding/json"
	"movie_api/helper"
	"movie_api/models/request"
	"movie_api/models/response"
	"net/http"
	"strconv"
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
	customerResponse, err := c.Services.GetAllUser(r.Context())
	helper.PrintError(err)

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(customerResponse)
	helper.PrintError(err)
}

func (c *Controller) GetUserById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	helper.PrintError(err)

	customerResponse, err := c.Services.GetUserById(r.Context(), idInt)
	helper.PrintError(err)

	if customerResponse != nil {
		w.Header().Set("content-type", "application/json")
		encoder := json.NewEncoder(w)
		err = encoder.Encode(customerResponse)
		helper.PrintError(err)
	} else {
		notFound := response.UserNotFound{
			Message: "User Not Found",
		}

		w.Header().Set("content-type", "application/json")
		encoder := json.NewEncoder(w)
		err = encoder.Encode(notFound)
		helper.PrintError(err)
	}
}

func (c *Controller) GetMovieList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) GetScreeningList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
