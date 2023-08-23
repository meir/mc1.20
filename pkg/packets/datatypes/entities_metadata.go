package datatypes

type EntityMetadata struct {
	Index uint8        `packet:"ubyte"`
	Type  MetadataType `packet:"varint"`
	Value interface{}  `packet:"metadata_value"`
}

type MetadataType int

const (
	MetadataTypeByte MetadataType = iota
	MetadataTypeVarInt
	MetadataTypeVarLong
	MetadataTypeFloat
	MetadataTypeString
	MetadataTypeChat
	MetadataTypeOptChat
	MetadataTypeSlot
	MetadataTypeBoolean
	MetadataTypeRotation
	MetadataTypePosition
	MetadataTypeOptPosition
	MetadataTypeDirection
	MetadataTypeOptUUID
	MetadataTypeBlockId
	MetadataTypeOptBlockId
	MetadataTypeNBT
	MetadataTypeParticle
	MetadataTypeVillagerData
	MetadataTypeOptVarInt
	MetadataTypePose
	MetadataTypeCatVariant
	MetadataTypeFrogVariant
	MetadataTypeOptGlobalPos
	MetadataTypePaintingVariant
	MetadataTypeSnifferState
	MetadataTypeVector3
	MetadataTypeQuaternion
)

type VillagerType int

const (
	VillagerTypeDesert VillagerType = iota
	VillagerTypeJungle
	VillagerTypePlains
	VillagerTypeSavanna
	VillagerTypeSnow
	VillagerTypeSwamp
	VillagerTypeTaiga
)

type VillagerProfession int

const (
	VillagerProfessionNone VillagerProfession = iota
	VillagerProfessionArmorer
	VillagerProfessionButcher
	VillagerProfessionCartographer
	VillagerProfessionCleric
	VillagerProfessionFarmer
	VillagerProfessionFisherman
	VillagerProfessionFletcher
	VillagerProfessionLeatherworker
	VillagerProfessionLibrarian
	VillagerProfessionMason
	VillagerProfessionNitwit
	VillagerProfessionShepherd
	VillagerProfessionToolSmith
	VillagerProfessionWeaponSmith
)

type Pose int

const (
	PoseStanding Pose = iota
	PoseFallFlying
	PoseSleeping
	PoseSwimming
	PoseSpinAttack
	PoseSneaking
	PoseLongJumping
	PoseDying
	PoseCroaking
	PoseUsingTongue
	PoseSitting
	PoseRoaring
	PoseSniffing
	PoseEmerging
	PoseDigging
)

type Direction int

const (
	DirectionDown Direction = iota
	DirectionUp
	DirectionNorth
	DirectionSouth
	DirectionWest
	DirectionEast
)

type SnifferState int

const (
	SnifferStateIdle SnifferState = iota
	SnifferStateHappy
	SnifferStateScenting
	SnifferStateSniffing
	SnifferStateSearching
	SnifferStateDigging
	SnifferStateRising
)

type DisplayType uint8

const (
	DisplayTypeNone DisplayType = iota
	DisplayTypeThirdPersonLeftHand
	DisplayTypeThirdPersonRightHand
	DisplayTypeFirstPersonLeftHand
	DisplayTypeFirstPersonRightHand
	DisplayTypeHead
	DisplayTypeGui
	DisplayTypeGround
	DisplayTypeFixed
)

type BoatType uint8

const (
	BoatTypeOak BoatType = iota
	BoatTypeSpruce
	BoatTypeBirch
	BoatTypeJungle
	BoatTypeAcacia
	BoatTypeDarkOak
)
