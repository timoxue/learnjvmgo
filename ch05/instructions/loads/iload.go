package loads

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//ILoad #
type ILoad struct{ base.Index8Instruction }

//ILoad0 #
type ILoad0 struct{ base.NoOperandsInstruction }

//ILoad1 #
type ILoad1 struct{ base.NoOperandsInstruction }

//ILoad2 #
type ILoad2 struct{ base.NoOperandsInstruction }

//ILoad3 #
type ILoad3 struct{ base.NoOperandsInstruction }

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

//Execute #
func (iload *ILoad) Execute(frame *rtda.Frame) {
	_iload(frame, uint(iload.Index))
}

//Execute #
func (iload1 *ILoad1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}
