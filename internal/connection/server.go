package connection

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net"

	"golang.org/x/exp/slog"
)

type Server struct {
	Host string
	Port string

	PrivateKey *rsa.PrivateKey
}

func NewServer(host, port string) (*Server, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, err
	}

	return &Server{
		Host: host,
		Port: port,

		PrivateKey: privateKey,
	}, nil
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

		go NewConnection(conn, s).Manage()
	}
}
