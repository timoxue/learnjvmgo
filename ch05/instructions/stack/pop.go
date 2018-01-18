package stack

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//Pop #
type Pop struct {
	base.NoOperandsInstruction
}

//Pop2 #
type Pop2 struct {
	base.NoOperandsInstruction
}

//Execute #
func (pop *Pop) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

//Execute #
func (op2 *Pop2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
