package constants

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//BIPush #
type BIPush struct{ val int8 }

//SIPush #
type SIPush struct{ val int16 }

//FetchOperands #
func (bipush *BIPush) FetchOperands(reader *base.BytecodeReader) {
	bipush.val = reader.ReadInt8()
}

//Execute #
func (bipush *BIPush) Execute(frame *rtda.Frame) {
	i := int32(bipush.val)
	frame.OperandStack().PushInt(i)
}
