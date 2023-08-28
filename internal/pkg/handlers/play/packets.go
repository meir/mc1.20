package play

import (
	"github.com/beito123/nbt"
	"github.com/meir/mc1.20/pkg/packets/datatypes"
)

type PacketLoginPlay struct {
	EntityID            int         `packet:"int"`
	IsHardcore          bool        `packet:"bool"`
	Gamemode            uint8       `packet:"uint8"`
	PreviousGamemode    uint8       `packet:"uint8"`
	DimensionNames      []string    `packet:"identifier,array"`
	RegistryCodec       *nbt.Stream `packet:"nbt"`
	DimensionType       string      `packet:"identifier"`
	DimensionName       string      `packet:"identifier"`
	HashedSeed          int64       `packet:"long"`
	MaxPlayers          int         `packet:"varint"`
	ViewDistance        int         `packet:"varint"`
	SimulationDistance  int         `packet:"varint"`
	ReducedDebugInfo    bool        `packet:"bool"`
	EnableRespawnScreen bool        `packet:"bool"`
	IsDebug             bool        `packet:"bool"`
	IsFlat              bool        `packet:"bool"`
	DeathLocation       struct {
		Dimension string             `packet:"identifier"`
		Location  datatypes.Position `packet:"position"`
	} `packet:"struct,optional"`
	PortalCooldown int `packet:"varint"`
}
