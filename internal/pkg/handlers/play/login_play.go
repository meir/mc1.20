package play

// import (
// 	"bufio"
// 	"fmt"

// 	"github.com/meir/mc1.20/internal/connection"
// 	"github.com/meir/mc1.20/pkg/packets"
// 	"golang.org/x/exp/slog"
// )

// func init() {
// 	connection.RegisterHandler(connection.StateLogin, connection.PacketId(connection), HandleLoginPlay)
// }

// func HandleLoginPlay(conn *connection.Connection, reader *bufio.Reader, packet packets.Packet) (bool, error) {
// 	loginStart := PacketEncryptionResponse{}
// 	err := packets.Unmarshal(reader, &loginStart)
// 	if err != nil {
// 		return false, err
// 	}

// 	slog.Debug(fmt.Sprintf("send success to %s", conn.Username))

// 	conn.Mutex.Lock()
// 	defer conn.Mutex.Unlock()
// 	conn.State = connection.StatePlay

// 	conn.Write(connection.PacketId(connection.ClientPacketLoginSuccess), PacketLoginSuccess{
// 		UUID:     conn.UUID,
// 		Username: conn.Username,
// 	})

// 	return true, nil
// }
