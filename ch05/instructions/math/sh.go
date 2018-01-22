package math

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//Ishl #
type Ishl struct{ base.NoOperandsInstruction }

//Ishr #
type Ishr struct{ base.NoOperandsInstruction }

//IUshr #
type IUshr struct{ base.NoOperandsInstruction }

//Lshl #
type Lshl struct{ base.NoOperandsInstruction }

//Lshr #
type Lshr struct{ base.NoOperandsInstruction }

//LUshr #
type LUshr struct{ base.NoOperandsInstruction }

//Execute #
func (ishl *Ishl) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

//Execute #
func (lshr *Lshr) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}
