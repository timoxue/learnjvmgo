package extended

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//GotoW #
type GotoW struct {
	offset int
}

//FetchOperands #
func (gotow *GotoW) FetchOperands(reader *base.BytecodeReader) {
	gotoW.offset = int(reader.ReadInt32())
}

//Execute #
func (gotow *GotoW) Execute(frame *rtda.Frame) {
	base.Branch(frame, gotow.offset)
}
