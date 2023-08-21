package main

import (
	"github.com/meir/mc1.20/internal/server"
	_ "github.com/meir/mc1.20/pkg/packets/parsers"
)

func main() {
	s := server.Server{
		Host: "0.0.0.0",
		Port: "25565",
	}

	err := s.Listen()
	if err != nil {
		panic(err)
	}
}
