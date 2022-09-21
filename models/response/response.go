package response

import "time"

type InsertUsers struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Users struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Movie struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LengthMin int    `json:"length_min"`
}

type MovieList struct {
	DataMovie interface{} `json:"data_movie"`
}

type ScreeningList struct {
	ID            int         `json:"id"`
	NameFilm      string      `json:"name_film"`
	ScreeningList interface{} `json:"screening_list"`
}

type ScreeningData struct {
	ID        int       `json:"id"`
	Room      string    `json:"room"`
	StartTime time.Time `json:"start_time"`
}

type UserNotFound struct {
	Message string `json:"message"`
}
