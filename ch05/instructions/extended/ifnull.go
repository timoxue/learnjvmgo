package extended

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//IFNull #
type IFNull struct {
	base.BranchInstruction
}

//IFNoNNull #
type IFNoNNull struct {
	base.BranchInstruction
}

//Execute #
func (ifnull *IFNull) Execute(frame *rtda.Frame) {
	ref := frame.OpreandStack().PopRef()
	if ref == nil {
		base.Branch(frame, ifnull.Offset)
	}
}
