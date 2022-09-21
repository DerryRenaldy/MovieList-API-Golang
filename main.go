package main

import "movie_api/server"

func main() {
	sv := server.Register()

	sv.Start()
}
