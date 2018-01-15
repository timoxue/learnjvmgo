package constants

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//NOP #
type NOP struct {
	base.NoOperandsInstruction
}

//Execute #
func (bcr *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}
