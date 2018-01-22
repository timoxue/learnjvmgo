package comparisons

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//FcmpG #
type FcmpG struct{ base.NoOperandsInstruction }

//FcmpL #
type FcmpL struct{ base.NoOperandsInstruction }

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

//Excute #
func (fcmpg *FcmpG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

//Execute #
func (fcmpl *FcmpL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
