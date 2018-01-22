package math

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
	"math"
)

//Drem #
type Drem struct{ base.NoOperandsInstruction }

//Frem #
type Frem struct{ base.NoOperandsInstruction }

//Irem #
type Irem struct{ base.NoOperandsInstruction }

//Lrem #
type Lrem struct{ base.NoOperandsInstruction }

//Execute #
func (irem *Irem) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArthmeticException : / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

//Execute #
func (drem *Drem) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
