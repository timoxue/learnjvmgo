package control

import "learnjvmgo/ch05/instructions/base"

type Goto struct{ base.BranchInstruction }

func (goto *Goto) Execute(frame *rtda.Frame) {
	base.Branch(frame, goto.Offset)
}
