package server

import (
	"fmt"
	"net"

	_ "github.com/meir/mc1.20/internal/pkg/handlers/handshake"
	_ "github.com/meir/mc1.20/internal/pkg/handlers/login"
	_ "github.com/meir/mc1.20/internal/pkg/handlers/play"
	_ "github.com/meir/mc1.20/internal/pkg/handlers/status"
	_ "github.com/meir/mc1.20/pkg/packets/parsers"

	"github.com/meir/mc1.20/internal/pkg/connection"
	"golang.org/x/exp/slog"
)

type Server struct {
	Host string
	Port string
}

func (s *Server) Listen() error {
	ip := fmt.Sprintf("%v:%v", s.Host, s.Port)

	host, err := net.Listen("tcp", ip)
	if err != nil {
		return err
	}

	slog.Info("starting server", "ip", ip)

	for {
		conn, err := host.Accept()
		if err != nil {
			return err
		}

		slog.Info("new connection", "ip", conn.RemoteAddr().String())

		go connection.NewConnection(conn).Manage()
	}
}
