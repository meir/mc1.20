package connection

import "github.com/meir/mc1.20/pkg/packets/datatypes"

type PacketLogout struct {
	Reason datatypes.Chat `packet:"chat"`
}

func (c Connection) DisconnectUser(onLogin bool, reason datatypes.Chat) error {
	packet := PacketLogout{
		Reason: reason,
	}

	id := PacketId(ClientPacketDisconnectLogin)
	if !onLogin {
		id = PacketId(ClientPacketDisconnectPlay)
	}

	err := c.Write(id, packet)
	if err != nil {
		return err
	}

	return nil
}
