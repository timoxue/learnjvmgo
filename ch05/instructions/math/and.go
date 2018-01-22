package math

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//Iand #
type Iand struct{ base.NoOperandsInstruction }

//Land #
type Land struct{ base.NoOperandsInstruction }

//Execute #
func (iland *Land) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

//FetchOperands #
func (iinc *IInc) FetchOperands(reader *base.BytecodeReader) {
	iinc.Index = uint(reader.ReadUint8())
	iinc.Const = int32(reader.ReadInt8())
}

//Execute #
func (iinc *IInc) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(iinc.Index)
	val += iinc.Const
	localVars.SetInt(iinc.Index, val)
}
