package connection

import (
	"bytes"

	"github.com/meir/mc1.20/pkg/packets"
)

type PacketHandler func(conn *Connection, reader *bytes.Reader, packet packets.Packet) (bool, error)

type ConnectionState int

const (
	StateHandshake ConnectionState = iota
	StateStatus
	StateLogin
	StatePlay
)

type Handlers []PacketHandler

func (h Handlers) Handle(conn *Connection, reader *bytes.Reader, packet packets.Packet) (bool, error) {
	for _, handler := range h {
		ok, err := handler(conn, reader, packet)
		if err != nil {
			return false, err
		}

		if ok {
			return true, nil
		}
	}

	return false, nil
}

var handlers map[ConnectionState]map[PacketId]Handlers

func RegisterHandler(state ConnectionState, event PacketId, handler PacketHandler) {
	if _, ok := handlers[state]; !ok {
		handlers[state] = map[PacketId]Handlers{}
	}

	if _, ok := handlers[state][event]; !ok {
		handlers[state][event] = []PacketHandler{}
	}

	handlers[state][event] = append([]PacketHandler{handler}, handlers[state][event]...)
}

func GetHandlers(state ConnectionState, event PacketId) Handlers {
	if stateMap, ok := handlers[state]; ok {
		if handlersArray, ok := stateMap[event]; ok {
			return handlersArray
		}
	}
	return []PacketHandler{}
}

type PacketId int

type HandshakePacketId PacketId
type StatusPacketId PacketId
type LoginPacketId PacketId
type PlayPacketId PacketId

const (
	ServerPacketHandshake       HandshakePacketId = iota
	ServerPacketLegacyHandshake HandshakePacketId = 0xFE
)

const (
	ClientPacketStatusResponse StatusPacketId = iota
	ClientPacketPingResponse
)

const (
	ServerPacketStatusRequest StatusPacketId = iota
	ServerPacketPingRequest
)

const (
	ClientPacketDisconnectLogin LoginPacketId = iota
	ClientPacketEncryptionRequest
	ClientPacketLoginSuccess
	ClientPacketSetCompression
	ClientPacketLoginPluginRequest
)

const (
	ServerPacketLoginStart LoginPacketId = iota
	ServerPacketEncryptionResponse
	ServerPacketLoginPluginResponse
)

const (
	ClientPacketBundleDelimiter PlayPacketId = iota
	ClientPacketSpawnEntity
	ClientPacketSpawnExperienceOrb
	ClientPacketSpawnPlayer
	ClientPacketEntityAnimation
	ClientPacketAwardStatistic
	ClientPacketAcknowledgeBlockChange
	ClientPacketSetBlockDestroyStage
	ClientPacketBlockEntityData
	ClientPacketBlockAction
	ClientPacketBlockUpdate
	ClientPacketBossBar
	ClientPacketChangeDifficulty
	ClientPacketChunkBiomes
	ClientPacketClearTitles
	ClientPacketCommandSuggestionsResponse
	ClientPacketCommands
	ClientPacketCloseContainer
	ClientPacketSetContainerContent
	ClientPacketSetContainerProperty
	ClientPacketSetContainerSlot
	ClientPacketSetCooldown
	ClientPacketChatSuggestions
	ClientPacketPluginMessage
	ClientPacketDamageEvent
	ClientPacketDeleteMessage
	ClientPacketDisconnectPlay
	ClientPacketDisguisedChatMessage
	ClientPacketEntityEvent
	ClientPacketExplosion
	ClientPacketUnloadChunk
	ClientPacketGameEvent
	ClientPacketOpenHorseScreen
	ClientPacketHurtAnimation
	ClientPacketInitializeWorldBorder
	ClientPacketKeepAlive
	ClientPacketChunkDataAndUpdateLight
	ClientPacketWorldEvent
	ClientPacketParticle
	ClientPacketUpdateLight
	ClientPacketLoginPlay
	ClientPacketMapData
	ClientPacketMerchantOffers
	ClientPacketUpdateEntityPosition
	ClientPacketUpdateEntityPositionAndRotation
	ClientPacketUpdateEntityRotation
	ClientPacketMoveVehicle
	ClientPacketOpenBook
	ClientPacketOpenScreen
	ClientPacketOpenSignEditor
	ClientPacketPingPlay
	ClientPacketPlaceGhostRecipe
	ClientPacketPlayerAbilities
	ClientPacketPlayerChatMessage
	ClientPacketEndCombat
	ClientPacketEnterCombat
	ClientPacketCombatDeath
	ClientPacketPlayerInfoRemove
	ClientPacketPlayerInfoUpdate
	ClientPacketLookAt
	ClientPacketSyncPlayerPosition
	ClientPacketUpdateRecipeBook
	ClientPacketRemoveEntities
	ClientPacketRemoveEntityEffect
	ClientPacketResourcePack
	ClientPacketRespawn
	ClientPacketSetHeadRotation
	ClientPacketUpdateSectionBlocks
	ClientPacketSelectAdvancementsTab
	ClientPacketServerData
	ClientPacketSetActionBarText
	ClientPacketSetBorderCenter
	ClientPacketSetBorderLerpSize
	ClientPacketSetBorderSize
	ClientPacketSetBorderWarningDelay
	ClientPacketSetBorderWarningDistance
	ClientPacketSetCamera
	ClientPacketSetHeldItem
	ClientPacketSetCenterChunk
	ClientPacketSetRenderDistance
	ClientPacketSetDefaultSpawnPosition
	ClientPacketDisplayObjective
	ClientPacketSetEntityMetadata
	ClientPacketLinkEntities
	ClientPacketSetEntityVelocity
	ClientPacketSetEquipment
	ClientPacketSetExperience
	ClientPacketSetHealth
	ClientPacketUpdateObjectives
	ClientPacketSetPassengers
	ClientPacketUpdateTeams
	ClientPacketUpdateScore
	ClientPacketSetSimulationDistance
	ClientPacketSetSubtitleText
	ClientPacketUpdateTime
	ClientPacketSetTitleText
	ClientPacketSetTitleAnimationTimes
	ClientPacketEntitySoundEffect
	ClientPacketSoundEffect
	ClientPacketStopSound
	ClientPacketSystemChatMessage
	ClientPacketSetTabListHeaderAndFooter
	ClientPacketTagQueryResponse
	ClientPacketPickupItem
	ClientPacketTeleportEntity
	ClientPacketUpdateAdvancements
	ClientPacketUpdateAttributes
	ClientPacketFeatureFlags
	ClientPacketEntityEffect
	ClientPacketUpdateRecipes
	ClientPacketUpdateTags
)
