package main

import (
	"os"

	"github.com/meir/mc1.20/internal/server"
	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	s := server.Server{
		Host: "0.0.0.0",
		Port: "25565",
	}

	err := s.Listen()
	if err != nil {
		panic(err)
	}
}
