package control

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//TableSwitch #
type TableSwitch struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

//FetchOperands #
func (tableswitch *TableSwitch) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	tableswitch.defaultOffset = reader.ReadInt32()
	tableswitch.low = reader.ReadInt32()
	tableswitch.high = reader.ReadInt32()
	jumpOffsetsCount := tableswitch.high - tableswitch.low + 1
	tableswitch.jumpOffsets = reader.ReadInt32s(jumpOffSetsCount)
}

//Execute #
func (tableswitch *TableSwitch) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= tableswitch.low && index <= tableswitch.high {
		offset = int(tableswitch.jumpOffsets[index-tableswitch.low])
	} else {
		offset = int(tableswitch.defaultOffset)
	}
	base.Branch(frame, offset)
}
