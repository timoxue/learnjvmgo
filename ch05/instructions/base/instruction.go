package base

import "learnjvmgo/ch05/rtda"

//Instruction #
type Instruction interface {
	FetchOprands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

//NoOperandsInstruction #
type NoOperandsInstruction struct{}

func (noi *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//nothing to do
}

//BranchInstruction #
type BranchInstruction struct {
	Offset int
}

func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	bi.Offset = int(reader.ReadInt16())
}
