package control

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

type LookupSwitch struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (lookupSwitch LookupSwitch) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	lookupSwitch.defaultOffset = reader.ReadInt32()
	lookupSwitch.npairs = reader.ReadInt32()
	lookupSwitch.matchOffsets = reader.ReadInt32s(lookupSwitch.npairs * 2)
}

func (lookupSwitch *LookupSwitch) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < lookupSwitch.npairs*2; i += 2 {
		if lookupSwitch.mathchOffsets[i] == key {
			offset := lookupSwitch.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(lookupSwitch.defaultOffset))
}
