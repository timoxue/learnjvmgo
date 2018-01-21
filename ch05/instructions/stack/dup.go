package stack

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//Dup #
type Dup struct{ base.NoOperandsInstruction }

//DupX1 #
type DupX1 struct{ base.NoOperandsInstruction }

//DupX2 #
type DupX2 struct{ base.NoOperandsInstruction }

//Dup2 #
type Dup2 struct{ base.NoOperandsInstruction }

//Dup2X1 #
type Dup2X1 struct{ base.NoOperandsInstruction }

//Dup2X2 #
type Dup2X2 struct{ base.NoOperandsInstruction }

func (dup *Dup) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
