package constants

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//AConstNULL #
type AConstNULL struct {
	base.NoOperandsInstruction
}

//DConst0 #
type DConst0 struct {
	base.NoOperandsInstruction
}

//DConst1 #
type DConst1 struct {
	base.NoOperandsInstruction
}

//FConst0 #
type FConst0 struct {
	base.NoOperandsInstruction
}

//FConst1 #
type FConst1 struct {
	base.NoOperandsInstruction
}

//FConst2 #
type FConst2 struct {
	base.NoOperandsInstruction
}

//IConstM1 #
type IConstM1 struct {
	base.NoOperandsInstruction
}

//IConst0 #
type IConst0 struct {
	base.NoOperandsInstruction
}

//IConst1 #
type IConst1 struct {
	base.NoOperandsInstruction
}

//IConst2 #
type IConst2 struct {
	base.NoOperandsInstruction
}

//IConst3 #
type IConst3 struct {
	base.NoOperandsInstruction
}

//IConst4 #
type IConst4 struct {
	base.NoOperandsInstruction
}

//IConst5 #
type IConst5 struct {
	base.NoOperandsInstruction
}

//LConst0 #
type LConst0 struct {
	base.NoOperandsInstruction
}

//LConst1 #
type LConst1 struct {
	base.NoOperandsInstruction
}

//Execute #
func (aConstNull *AConstNULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

//Execute #
func (dConst0 *DConst0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

//Execute #
func (iConstM1 *IConstM1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}
