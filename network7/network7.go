package network7

const (
	MaxClients    = 64
	NetVersion    = "0.7 802f1be60a05665f"
	ClientVersion = 0x0705

	ChatAll     ChatMode = 1
	ChatTeam    ChatMode = 2
	ChatWhisper ChatMode = 3

	TeamSpectators GameTeam = -1
	TeamRed        GameTeam = 0
	TeamBlue       GameTeam = 1

	VoteUnknown   Vote = 0
	VoteStartOp   Vote = 1
	VoteStartKick Vote = 2
	VoteStartSpec Vote = 3
	VoteEndAbort  Vote = 4
	VoteEndPass   Vote = 5
	VoteEndFail   Vote = 6

	// oop!
	EmoteOop Emote = 0
	// !
	EmoteExclamation Emote = 1
	EmoteHearts      Emote = 2
	// tear
	EmoteDrop Emote = 3
	// ...
	EmoteDotdot Emote = 4
	EmoteMusic  Emote = 5
	EmoteSorry  Emote = 6
	EmoteGhost  Emote = 7
	// annoyed
	EmoteSushi Emote = 8
	// angry
	EmoteSplattee Emote = 9
	EmoteDeviltee Emote = 10
	// swearing
	EmoteZomg Emote = 11
	EmoteZzz  Emote = 12
	EmoteWtf  Emote = 13
	// happy
	EmoteEyes Emote = 14
	// ??
	EmoteQuestion Emote = 15

	MsgCtrlKeepAlive = 0x00
	MsgCtrlConnect   = 0x01
	MsgCtrlAccept    = 0x02
	MsgCtrlToken     = 0x05
	MsgCtrlClose     = 0x04

	// TODO: these should preferrably all be devide dinto different type dintegers
	// same as ChatMode, etc. so that the user can easily see which integer to pass
	// to which function as which parameter
	MsgSysInfo            = 1
	MsgSysMapChange       = 2
	MsgSysMapData         = 3
	MsgSysServerInfo      = 4
	MsgSysConReady        = 5
	MsgSysSnap            = 6
	MsgSysSnapEmpty       = 7
	MsgSysSnapSingle      = 8
	MsgSysSnapSmall       = 9
	MsgSysInputTiming     = 10
	MsgSysRconAuthOn      = 11
	MsgSysRconAuthOff     = 12
	MsgSysRconLine        = 13
	MsgSysRconCmdAdd      = 14
	MsgSysRconCmdRem      = 15
	MsgSysAuthChallenge   = 16 // unused
	MsgSysAuthResult      = 17 // unused
	MsgSysReady           = 18
	MsgSysEnterGame       = 19
	MsgSysInput           = 20
	MsgSysRconCmd         = 21
	MsgSysRconAuth        = 22
	MsgSysRequestMapData  = 23
	MsgSysAuthStart       = 24 // unused
	MsgSysAuthResponse    = 25 // unused
	MsgSysPing            = 26
	MsgSysPingReply       = 27
	MsgSysError           = 28 // unused
	MsgSysMaplistEntryAdd = 29
	MsgSysMaplistEntryRem = 30

	MsgGameSvMotd              = 1
	MsgGameSvBroadcast         = 2
	MsgGameSvChat              = 3
	MsgGameSvTeam              = 4
	MsgGameSvKillMsg           = 5
	MsgGameSvTuneParams        = 6
	MsgGameSvExtraProjectile   = 7 // unused
	MsgGameSvReadyToEnter      = 8
	MsgGameSvWeaponPickup      = 9
	MsgGameSvEmoticon          = 10
	MsgGameSvVoteClearOptions  = 11
	MsgGameSvVoteOptionListAdd = 12
	MsgGameSvVoteOptionAdd     = 13
	MsgGameSvVoteOptionRemove  = 14
	MsgGameSvVoteSet           = 15
	MsgGameSvVoteStatus        = 16
	MsgGameSvServerSettings    = 17
	MsgGameSvClientInfo        = 18
	MsgGameSvGameInfo          = 19
	MsgGameSvClientDrop        = 20
	MsgGameSvGameMsg           = 21
	MsgGameDeClientEnter       = 22
	MsgGameDeClientLeave       = 23
	MsgGameClSay               = 24
	MsgGameClSetTeam           = 25
	MsgGameClSetSpectatorMode  = 26
	MsgGameClStartInfo         = 27
	MsgGameClKill              = 28
	MsgGameClReadyChange       = 29
	MsgGameClEmoticon          = 30
	MsgGameClVote              = 31
	MsgGameClCallVote          = 32
	MsgGameSvSkinChange        = 33
	MsgGameClSkinChange        = 34
	MsgGameSvRaceFinish        = 35
	MsgGameSvCheckpoint        = 36
	MsgGameSvCommandInfo       = 37
	MsgGameSvCommandInfoRemove = 38
	MsgGameClCommand           = 39

	TypeControl  MsgType = 1
	TypeNet      MsgType = 2
	TypeConnless MsgType = 3

	// can be sent by the server in kill messages
	WeaponGame  Weapon = -3
	WeaponSelf  Weapon = -2
	WeaponWorld Weapon = -1

	// can be sent by the client when requesting weapon switch
	// or by the server in kill messages
	WeaponHammer  Weapon = 0
	WeaponGun     Weapon = 1
	WeaponShotgun Weapon = 2
	WeaponGrenade Weapon = 3
	WeaponLaser   Weapon = 4
	WeaponNinja   Weapon = 5
	NumWeapons    Weapon = 6
)

type Vote int
type Emote int
type ChatMode int
type GameTeam int
type Weapon int

type MsgType int
