package main

import "goreservationapp/pkg/server"

func main() {
	s := server.NewServer()
	s.Start()
}
