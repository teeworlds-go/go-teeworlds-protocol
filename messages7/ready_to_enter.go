package messages7

import (
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type ReadyToEnter struct {
}

func (msg ReadyToEnter) MsgId() int {
	return network7.MsgGameReadyToEnter
}

func (msg ReadyToEnter) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg ReadyToEnter) System() bool {
	return false
}

func (msg ReadyToEnter) Vital() bool {
	return true
}

func (msg ReadyToEnter) Pack() []byte {
	return []byte{}
}

func (msg *ReadyToEnter) Unpack(u *packer.Unpacker) {
}
