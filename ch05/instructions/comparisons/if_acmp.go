package comparisons

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//IFACMPeq #
type IFACMPeq struct{ base.BranchInstruction }

//IFACMPne #
type IFACMPne struct{ base.BranchInstruction }

//Execute #
func (eq *IFACMPeq) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, eq.Offset)
	}
}
