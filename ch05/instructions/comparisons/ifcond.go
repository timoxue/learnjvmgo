package comparisons

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//IFeq #
type IFeq struct{ base.BranchInstruction }

//IFne #
type IFne struct{ base.BranchInstruction }

//IFlt #
type IFlt struct{ base.BranchInstruction }

//IFle #
type IFle struct{ base.BranchInstruction }

//IFgt #
type IFgt struct{ base.BranchInstruction }

//IFge #
type IFge struct{ base.BranchInstruction }

//Execute #
func (ifeq *IFeq) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, ifeq.Offset)
	}
}
