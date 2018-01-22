package conversions

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//D2F #
type D2F struct{ base.NoOperandsInstruction }

//D2I #
type D2I struct{ base.NoOperandsInstruction }

//D2L #
type D2L struct{ base.NoOperandsInstruction }

//Execute #
func (d2i *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}
