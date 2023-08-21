package objects

type Entity struct {
	State         uint8 `packet:"byte"`
	AirTicks      int   `packet:"varint"`
	CustomName    *Chat `packet:"chat,optional"`
	IsNameVisible bool  `packet:"bool"`
	IsSilent      bool  `packet:"bool"`
	HasNoGravity  bool  `packet:"bool"`
	Pose          Pose  `packet:"pose"`
	FrostTicks    int   `packet:"varint"`
}

type EntityInteraction struct {
	*Entity

	Width      float32 `packet:"float"`
	Height     float32 `packet:"float"`
	Responsive bool    `packet:"bool"`
}

type EntityDisplay struct {
	*Entity

	InterpolationDelay    int        `packet:"varint"`
	InterpolationDuration int        `packet:"varint"`
	Translation           Vector3    `packet:"vector"`
	Scale                 Vector3    `packet:"vector"`
	RotationLeft          Quaternion `packet:"rotation"`
	RotationRight         Quaternion `packet:"rotation"`
	BillboardConstraint   uint8      `packet:"byte"`
	BrightnessOverride    int        `packet:"varint"`
	ViewRange             float32    `packet:"float"`
	ShadowRadius          float32    `packet:"float"`
	ShadowStrength        float32    `packet:"float"`
	Width                 float32    `packet:"float"`
	Height                float32    `packet:"float"`
	GlowColorOverride     int        `packet:"varint"`
}

type EntityBlockDisplay struct {
	*EntityDisplay

	BlockId int `packet:"varint"`
}

type EntityItemDisplay struct {
	*EntityDisplay

	Item interface{} `packet:"slot"` //TODO: type "Slot"
	Type uint8       `packet:"byte"`
}

type EntityTextDisplay struct {
	*EntityDisplay

	Text            Chat  `packet:"chat"`
	LineWidth       int   `packet:"varint"`
	BackgroundColor int   `packet:"varint"`
	TextOpacity     uint8 `packet:"byte"`
	Opts            uint8 `packet:"byte"`
}

type EntityThrownItemProjectile struct {
	*Entity

	Item interface{} `packet:"slot"` //TODO: type "Slot"
}

type EntityThrownEgg struct {
	*EntityThrownItemProjectile
}

type EntityThrownEnderPearl struct {
	*EntityThrownItemProjectile
}

type EntityThrownExperienceBottle struct {
	*EntityThrownItemProjectile
}

type EntityThrownPotion struct {
	*EntityThrownItemProjectile
}

type EntityThrownSnowball struct {
	*EntityThrownItemProjectile
}

type EntityEyeOfEnder struct {
	*Entity

	Item interface{} `packet:"slot"` //TODO: type "Slot"
}

type EntityFallingBlock struct {
	*Entity

	BlockPosition Position `packet:"position"`
}

type EntityAreaEffectCloud struct {
	*Entity

	Radius       float32     `packet:"float"`
	Color        int         `packet:"varint"`
	IgnoreRadius bool        `packet:"bool"`
	Particle     interface{} `packet:"particle"` //TODO: type "Particle"
}

type EntityFishingHook struct {
	*Entity

	HookId      int  `packet:"varint"`
	IsCatchable bool `packet:"bool"`
}

type EntityAbstractArrow struct {
	*Entity

	Opts     uint8 `packet:"byte"`
	Piercing uint8 `packet:"byte"`
}

type EntityArrow struct {
	*EntityAbstractArrow

	Color int `packet:"varint"`
}

type EntitySpectralArrow struct {
	*EntityAbstractArrow
}

type EntityThrownTrident struct {
	*EntityAbstractArrow

	LoyaltyLevel   int  `packet:"varint"`
	HasEnchantment bool `packet:"bool"`
}

type EntityBoat struct {
	*Entity

	HitTime          int      `packet:"varint"`
	ForwardDirection int      `packet:"varint"`
	Damagetaken      float32  `packet:"float"`
	Type             BoatType `packet:"varint"`
	IsLeftPaddling   bool     `packet:"bool"`
	IsRightPaddling  bool     `packet:"bool"`
	SplashTimer      int      `packet:"varint"`
}

type EntityChestBoat struct {
	*EntityBoat
}

type EntityEndCrystal struct {
	*Entity

	BeamTarget Position `packet:"position"`
	ShowBottom bool     `packet:"bool"`
}

type EntityDragonFireball struct {
	*Entity
}

type EntitySmallFireball struct {
	*Entity

	Item interface{} `packet:"slot"` //TODO: type "Slot"
}

type EntityFireball struct {
	*Entity

	Item interface{} `packet:"slot"` //TODO: type "Slot"
}

type EntityWitherSkull struct {
	*Entity

	IsInvulnerable bool `packet:"bool"`
}

type EntityFireworkRocket struct {
	*Entity

	Item        interface{} `packet:"slot"` //TODO: type "Slot"
	OptVarint   *int        `packet:"varint,optional"`
	ShotAtAngle bool        `packet:"bool"`
}

type EntityItemFrame struct {
	*Entity

	Item     interface{} `packet:"slot"` //TODO: type "Slot"
	Rotation int         `packet:"varint"`
}

type EntityGlowingItemFrame struct {
	*EntityItemFrame
}

type EntityPainting struct {
	*Entity

	PaintingVariant int `packet:"varint"` //TODO: type "PaintingVariant"
}

type EntityItem struct {
	*Entity

	Item interface{} `packet:"slot"` //TODO: type "Slot"
}

type EntityLiving struct {
	*Entity

	HandStates            uint8     `packet:"byte"`
	Health                float32   `packet:"float"`
	PotionEffectColor     int       `packet:"varint"`
	IsPotionEffectAmbient bool      `packet:"bool"`
	ArrowsStuck           int       `packet:"varint"`
	BeeStingersStuck      int       `packet:"varint"`
	BedPosition           *Position `packet:"position,optional"`
}

type EntityPlayer struct {
	*EntityLiving

	AdditionalHearts    float32     `packet:"float"`
	Score               int         `packet:"varint"`
	DisplayedSkinParts  uint8       `packet:"byte"`
	MainHand            uint8       `packet:"byte"`
	LeftShoulderEntity  interface{} `packet:"nbt,optional"` //TODO: type "NBT"
	RightShoulderEntity interface{} `packet:"nbt,optional"` //TODO: type "NBT"
}

type EntityArmorStand struct {
	*EntityLiving

	Parts            uint8   `packet:"byte"`
	HeadRotation     Vector3 `packet:"vector"`
	BodyRotation     Vector3 `packet:"vector"`
	LeftArmRotation  Vector3 `packet:"vector"`
	RightArmRotation Vector3 `packet:"vector"`
	LeftLegRotation  Vector3 `packet:"vector"`
	RightLegRotation Vector3 `packet:"vector"`
}

type EntityMob struct {
	*EntityLiving

	Opts uint8 `packet:"byte"`
}

type EntityAmbientCreature struct {
	*EntityMob
}

type EntityBat struct {
	*EntityAmbientCreature

	IsHanging bool `packet:"bool"`
}

type EntityPathfinderMob struct {
	*EntityMob
}

type EntityWaterAnimal struct {
	*EntityPathfinderMob
}

type EntitySquid struct {
	*EntityWaterAnimal
}

type EntityDolphin struct {
	*EntityWaterAnimal

	TreasurePosition Position `packet:"position"`
	HasFish          bool     `packet:"bool"`
	MoistureLevel    int      `packet:"varint"`
}

type EntityAbstractFish struct {
	*EntityWaterAnimal

	FromBucket bool `packet:"bool"`
}

type EntityCod struct {
	*EntityAbstractFish
}

type EntityPufferfish struct {
	*EntityAbstractFish

	PuffState int `packet:"varint"`
}

type EntitySalmon struct {
	*EntityAbstractFish
}

type EntityTropicalFish struct {
	*EntityAbstractFish

	Variant int `packet:"varint"`
}

type EntityTadpole struct {
	*EntityAbstractFish
}

type EntityAgeableMob struct {
	*EntityPathfinderMob

	IsBaby bool `packet:"bool"`
}

type EntityAnimal struct {
	*EntityAgeableMob
}

type EntitySniffer struct {
	*EntityAnimal

	SnifferState SnifferState `packet:"varint"`
	DropSeedTick int          `packet:"varint"`
}

type EntityAbstractHorse struct {
	*EntityAnimal

	Opts uint8 `packet:"byte"`
}

type EntityHorse struct {
	*EntityAbstractHorse

	Variant int `packet:"varint"`
}

type EntityZombieHorse struct {
	*EntityAbstractHorse
}

type EntitySkeletonHorse struct {
	*EntityAbstractHorse
}

type EntityCamel struct {
	*EntityAbstractHorse

	IsDashing    bool  `packet:"bool"`
	LastPoseTick int64 `packet:"long"`
}

type EntityChestedHorse struct {
	*EntityAbstractHorse

	HasChest bool `packet:"bool"`
}

type EntityDonkey struct {
	*EntityChestedHorse
}

type EntityLlama struct {
	*EntityChestedHorse

	Strength    int `packet:"varint"`
	CarpetColor int `packet:"varint"`
	Variant     int `packet:"varint"`
}

type EntityTraderLlama struct {
	*EntityLlama
}

type EntityMule struct {
	*EntityChestedHorse
}

type EntityAxolotl struct {
	*EntityAnimal

	Variant           int  `packet:"varint"`
	PlayingDead       bool `packet:"bool"`
	SpawnedFromBucket bool `packet:"bool"`
}

type EntityBee struct {
	*EntityAnimal

	Flags uint8 `packet:"byte"`
	Anger int   `packet:"varint"`
}

type EntityFox struct {
	*EntityAnimal

	Type    int          `packet:"varint"`
	Flags   uint8        `packet:"byte"`
	OptUUID *interface{} `packet:"uuid,optional"` //TODO: type "UUID"
	optUUID *interface{} `packet:"uuid,optional"` //TODO: type "UUID"
}

type EntityFrog struct {
	*EntityAnimal

	Variant   int `packet:"varint"`
	OptTarget int `packet:"varint,optional"`
}

type EntityOcelot struct {
	*EntityAnimal

	IsTrusting bool `packet:"bool"`
}

type EntityPanda struct {
	*EntityAnimal

	BreedTime  int   `packet:"varint"`
	SneezeTime int   `packet:"varint"`
	EatTime    int   `packet:"varint"`
	MainGene   int   `packet:"varint"`
	HiddenGene int   `packet:"varint"`
	Flags      uint8 `packet:"byte"`
}

type EntityPig struct {
	*EntityAnimal

	Saddle    bool `packet:"bool"`
	BoostTime int  `packet:"varint"`
}

type EntityRabbit struct {
	*EntityAnimal

	Type int `packet:"varint"`
}

type EntityTurtle struct {
	*EntityAnimal

	HomePosition   Position `packet:"position"`
	HasEgg         bool     `packet:"bool"`
	LayingEgg      bool     `packet:"bool"`
	TravelPosition Position `packet:"position"`
	GoingHome      bool     `packet:"bool"`
	Traveling      bool     `packet:"bool"`
}

type EntityPolarBear struct {
	*EntityAnimal

	StandingUp bool `packet:"bool"`
}

type EntityChicken struct {
	*EntityAnimal
}

type EntityCow struct {
	*EntityAnimal
}

type EntityHoglin struct {
	*EntityAnimal

	IsImmuneToZombification bool `packet:"bool"`
}

type EntityMooshroom struct {
	*EntityCow

	Type int `packet:"varint"`
}

type EntitySheep struct {
	*EntityAnimal

	Color int `packet:"varint"`
}

type EntityStrider struct {
	*EntityAnimal

	BoostTime int  `packet:"varint"`
	Shaking   bool `packet:"bool"`
	Saddle    bool `packet:"bool"`
}

type EntityTameableAnimal struct {
	*EntityAnimal

	Flags  uint8       `packet:"byte"`
	Ownder interface{} `packet:"uuid,optional"` //TODO: type "UUID"
}

type EntityCat struct {
	*EntityTameableAnimal

	Variant     int  `packet:"varint"`
	IsLying     bool `packet:"bool"`
	IsRelaxed   bool `packet:"bool"`
	CollarColor int  `packet:"varint"`
}

type EntityWolf struct {
	*EntityTameableAnimal

	IsBegging   bool `packet:"bool"`
	CollarColor int  `packet:"varint"`
	AngerTime   int  `packet:"varint"`
}

type EntityParrot struct {
	*EntityTameableAnimal

	Variant int `packet:"varint"`
}

type EntityAbstractVillager struct {
	*EntityMob

	HeadShakeTime int `packet:"varint"`
}

type EntityVillager struct {
	*EntityAbstractVillager

	VillagerData interface{} `packet:"nbt"` //TODO: type "VillagerData"
}

type EntityWanderingTrader struct {
	*EntityAbstractVillager
}

type EntityAbstractGolem struct {
	*EntityPathfinderMob
}

type EntityIronGolem struct {
	*EntityAbstractGolem

	Flags uint8 `packet:"byte"`
}

type EntitySnowGolem struct {
	*EntityAbstractGolem

	Flags uint8 `packet:"byte"`
}

type EntityShulker struct {
	*EntityMob

	AttachFace   Direction `packet:"varint"`
	OptPosition  *Position `packet:"position,optional"`
	ShieldHeight byte      `packet:"byte"`
	Color        uint8     `packet:"byte"`
}

type EntityMonster struct {
	*EntityPathfinderMob
}

type EntityBasePiglin struct {
	*EntityMonster

	IsImmuneToZombification bool `packet:"bool"`
}

type EntityPiglin struct {
	*EntityBasePiglin

	IsBaby             bool `packet:"bool"`
	IsChargingCrossbow bool `packet:"bool"`
	IsDancing          bool `packet:"bool"`
}

type EntityPiglinBrute struct {
	*EntityBasePiglin
}

type EntityBlaze struct {
	*EntityMonster

	Flags byte `packet:"byte"`
}

type EntityCreeper struct {
	*EntityMonster

	State     int  `packet:"varint"`
	IsCharged bool `packet:"bool"`
	IsIgnited bool `packet:"bool"`
}

type EntityEndermite struct {
	*EntityMonster
}

type EntityGiant struct {
	*EntityMonster
}

type EntityGoat struct {
	*EntityAnimal
}

type EntityGuardian struct {
	*EntityMonster

	IsRetractingSpikes bool `packet:"bool"`
	TargetEID          int  `packet:"varint"`
}

type EntityElderGuardian struct {
	*EntityGuardian
}

type EntitySilverfish struct {
	*EntityMonster
}

type EntityRaider struct {
	*EntityMonster

	IsCelebrating bool `packet:"bool"`
}

type EntityAbstractIllager struct {
	*EntityRaider
}

type EntityVindicator struct {
	*EntityAbstractIllager
}

type EntityPillager struct {
	*EntityAbstractIllager

	IsCharging bool `packet:"bool"`
}

type EntitySpellcasterIllager struct {
	*EntityAbstractIllager

	Spell uint8 `packet:"byte"`
}

type EntityEvoker struct {
	*EntitySpellcasterIllager
}

type EntityIllusioner struct {
	*EntitySpellcasterIllager
}

type EntityRavager struct {
	*EntityRaider
}

type EntityWitch struct {
	*EntityRaider

	IsDrinkingPotion bool `packet:"bool"`
}

type EntityEvokerFangs struct {
	*Entity
}

type EntityVex struct {
	*EntityMonster

	Flags byte `packet:"byte"`
}

type EntityAbstractSkeleton struct {
	*EntityMonster
}

type EntitySkeleton struct {
	*EntityAbstractSkeleton
}

type EntityWitherSkeleton struct {
	*EntityAbstractSkeleton
}

type EntityStray struct {
	*EntityAbstractSkeleton
}

type EntitySpider struct {
	*EntityMonster

	Flags byte `packet:"byte"`
}

type EntityWarden struct {
	*EntityMonster

	AngerLevel int `packet:"varint"`
}

type EntityWither struct {
	*EntityMonster

	CenterHeadTarget int `packet:"varint"`
	LeftHeadTarget   int `packet:"varint"`
	RightHeadTarget  int `packet:"varint"`
	InvulnerableTime int `packet:"varint"`
}

type EntityZoglin struct {
	*EntityMonster

	IsBaby bool `packet:"bool"`
}

type EntityZombie struct {
	*EntityMonster

	IsBaby          bool `packet:"bool"`
	Type            int  `packet:"varint"` // deprecated
	BecomingDrowned bool `packet:"bool"`
}

type EntityZombieVillager struct {
	*EntityZombie

	IsConverting bool        `packet:"bool"`
	VillagerData interface{} `packet:"nbt"` //TODO: type "VillagerData"
}

type EntityHusk struct {
	*EntityZombie
}

type EntityDrowned struct {
	*EntityZombie
}

type EntityZombifiedPiglin struct {
	*EntityZombie
}

type EntityEnderman struct {
	*EntityMonster

	OptBlockId  *int `packet:"varint,optional"` //TODO: type "BlockID"
	IsScreaming bool `packet:"bool"`
	IsStaring   bool `packet:"bool"`
}

type EntityEnderDragon struct {
	*EntityMonster

	DragonPhase int `packet:"varint"`
}

type EntityFlying struct {
	*EntityMob
}

type EntityGhast struct {
	*EntityFlying

	IsAttacking bool `packet:"bool"`
}

type EntityPhantom struct {
	*EntityFlying

	Size int `packet:"varint"`
}

type EntitySlime struct {
	*EntityMob

	Size int `packet:"varint"`
}

type EntityLlamaSpit struct {
	*Entity
}

type EntityAbstractMinecart struct {
	*Entity

	ShakingPower      int     `packet:"varint"`
	ShakingDirection  int     `packet:"varint"`
	ShakingMultiplier float32 `packet:"float"`
	BlockId           int     `packet:"varint"`
	BlockOffset       int     `packet:"varint"`
	ShowBlock         bool    `packet:"bool"`
}

type EntityMinecart struct {
	*EntityAbstractMinecart
}

type EntityAbstractMinecartContainer struct {
	*EntityAbstractMinecart
}

type EntityMinecartHopper struct {
	*EntityAbstractMinecartContainer
}

type EntityMinecartChest struct {
	*EntityAbstractMinecartContainer
}

type EntityMinecartFurnace struct {
	*EntityAbstractMinecart

	HasFuel bool `packet:"bool"`
}

type EntityMinecartTNT struct {
	*EntityAbstractMinecart
}

type EntityAbstractMinecartSpawner struct {
	*EntityAbstractMinecart
}

type EntityMinecartCommandBlock struct {
	*EntityAbstractMinecart

	Command    string `packet:"string"`
	LastOutput Chat   `packet:"chat"`
}

type EntityPrimedTnt struct {
	*Entity

	FuseTime int `packet:"varint"`
}
