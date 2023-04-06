package main

import (
	"challenges_9/server"

	_ "github.com/lib/pq"
)

func main() {
	//run server

	server.NewServer()
}
