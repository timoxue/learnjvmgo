package stack

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//Swap #
type Swap struct{ base.NoOperandsInstruction }

//Execute #
func (swap *Swap) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
