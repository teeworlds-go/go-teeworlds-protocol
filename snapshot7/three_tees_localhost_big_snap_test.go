package snapshot7_test

// same connection as three_tees_localhost_test.go

import (
	"testing"

	"github.com/teeworlds-go/protocol/internal/testutils/require"
	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/object7"
	"github.com/teeworlds-go/protocol/protocol7"
)

// localhost first connection
// 2 players already connected
// map bridge_pickups

func TestBigSnapDuringEstablishedConnection(t *testing.T) {
	t.Parallel()
	// snapshot captured with tcpdump
	// generated by a vanilla teeworlds 0.7.5 server
	// used https://github.com/ChillerDragon/teeworlds/tree/hacking-on-protocol client to connect
	// 0.7 vanilla based client with debug prints
	//
	// libtw2 dissector details
	// Teeworlds 0.7 Protocol packet
	//     Flags: compressed (..01 00..)
	//     Acknowledged sequence number: 6 (.... ..00 0000 0110)
	//     Number of chunks: 1
	//     Token: 57d3edf3
	//     Compressed payload (41 bytes)
	// Teeworlds 0.7 Protocol chunk: sys.snap_single
	//     Header (non-vital)
	//     Message: sys.snap_single
	//     Tick: 1126
	//     Delta tick: 4
	//     Crc: 18606
	//     Data (54 bytes)
	dump := []byte{
		0x10, 0x06, 0x01, 0x57, 0xd3, 0xed, 0xf3,
		0x4d, 0x8b, 0x29, 0x4b, 0xa6, 0x3c, 0x6a,
		0x3e, 0x0b, 0x5f, 0x53, 0xd4, 0xa3, 0xe7,
		0xcf, 0xb8, 0xef, 0x4d, 0x80, 0xfc, 0xfe,
		0xbc, 0xff, 0xeb, 0x85, 0xc7, 0x97, 0x2d,
		0x36, 0x0d, 0x68, 0xae, 0xd2, 0xdb, 0xdf,
		0xf3, 0xe7, 0xfd, 0x5f, 0x71, 0x03,
	}

	packet := protocol7.Packet{}
	err := packet.Unpack(dump)
	require.NoError(t, err)

	// TODO: not working yet
	// conn := protocol7.Session{}
	// conn.Ack = packet.Header.Ack
	// repack := packet.Pack(&conn)
	// require.Equal(t, dump, repack)

	// content
	require.Equal(t, 1, len(packet.Messages))
	require.Equal(t, network7.MsgSysSnapSingle, packet.Messages[0].MsgId())
	msg, ok := packet.Messages[0].(*messages7.SnapSingle)
	require.Equal(t, true, ok)

	// verified with hacking on protocol print
	require.Equal(t, 2, msg.Snapshot.NumItemDeltas)
	require.Equal(t, 0, msg.Snapshot.NumRemovedItems)
	require.Equal(t, 2, len(msg.Snapshot.Items))

	// TODO: this should be 18606
	//       the snap only contains 2 items
	//       and they fully match what the hacking on protocol branch
	//       printed in the C++ version
	//       but the snap crc is not for the new delta items
	//       but for the full snapshot
	//       so we need all prior snapshots to compute the correct crc
	require.Equal(t, -1637, msg.Snapshot.Crc)

	// verified with hacking on protocol
	item := msg.Snapshot.Items[0]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok := item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 1, character.Id())
	require.Equal(t, 4, character.Tick)
	require.Equal(t, 11, character.X)
	require.Equal(t, -4, character.Y)
	require.Equal(t, -18, character.VelX)
	require.Equal(t, -558, character.VelY)
	require.Equal(t, 0, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, 0, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 11, character.HookX)
	require.Equal(t, -2, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, network7.WeaponHammer, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[1]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 2, character.Id())
	require.Equal(t, 4, character.Tick)
	require.Equal(t, -9, character.X)
	require.Equal(t, -35, character.Y)
	require.Equal(t, -1750, character.VelX)
	require.Equal(t, 691, character.VelY)
	require.Equal(t, 5, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, 0, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 4, character.HookTick)
	require.Equal(t, 11, character.HookX)
	require.Equal(t, -2, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, network7.WeaponHammer, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)
}
