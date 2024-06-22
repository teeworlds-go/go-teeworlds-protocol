package messages7

import (
	"slices"

	"github.com/teeworlds-go/teeworlds/chunk7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type SnapEmpty struct {
	ChunkHeader *chunk7.ChunkHeader

	GameTick  int
	DeltaTick int
}

func (msg SnapEmpty) MsgId() int {
	return network7.MsgSysSnapEmpty
}

func (msg SnapEmpty) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg SnapEmpty) System() bool {
	return true
}

func (msg SnapEmpty) Vital() bool {
	return false
}

func (msg SnapEmpty) Pack() []byte {
	return slices.Concat(
		packer.PackInt(msg.GameTick),
		packer.PackInt(msg.DeltaTick),
	)
}

func (msg *SnapEmpty) Unpack(u *packer.Unpacker) {
	msg.GameTick = u.GetInt()
	msg.DeltaTick = u.GetInt()
}

func (msg *SnapEmpty) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SnapEmpty) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
