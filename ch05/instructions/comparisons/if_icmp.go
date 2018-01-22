package comparisons

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//IFICMPeq #
type IFICMPeq struct{ base.BranchInstruction }

//IFICMPne #
type IFICMPne struct{ base.BranchInstruction }

//IFICMPlt #
type IFICMPlt struct{ base.BranchInstruction }

//IFICMPle #
type IFICMPle struct{ base.BranchInstruction }

//IFICMPgt #
type IFICMPgt struct{ base.BranchInstuction }

//IFICMPge #
type IFICMPge struct{ base.BranchInstruction }

//Execute #
func (ne *IFICMPne) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, ne.Offset)
	}
}
