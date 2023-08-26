package main

import (
	"os"

	_ "github.com/meir/mc1.20/internal/pkg/handlers/handshake"
	_ "github.com/meir/mc1.20/internal/pkg/handlers/login"
	_ "github.com/meir/mc1.20/internal/pkg/handlers/play"
	_ "github.com/meir/mc1.20/internal/pkg/handlers/status"
	_ "github.com/meir/mc1.20/pkg/packets/parsers"

	"github.com/meir/mc1.20/internal/connection"
	"golang.org/x/exp/slog"
)

func main() {
	// set up logging for debugging
	// TODO: make this configurable
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	// TODO: make host and port configurable
	s, err := connection.NewServer(
		"0.0.0.0",
		"25565",
	)
	if err != nil {
		panic(err)
	}

	err = s.Listen()
	if err != nil {
		panic(err)
	}
}
