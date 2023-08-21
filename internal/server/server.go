package server

import (
	"fmt"
	"net"

	"github.com/meir/mc1.20/pkg/connection"
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

		go connection.NewConnection(conn)
	}
}
