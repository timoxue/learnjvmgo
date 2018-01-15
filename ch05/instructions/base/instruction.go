package base

import "learnjvmgo/ch05/rtda"

//Instruction #
type Instruction interface {
	FetchOprands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

//NoOperandsInstruction #
type NoOperandsInstruction struct{}

//FetchOperands #
func (noi *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//nothing to do
}

//BranchInstruction #
type BranchInstruction struct {
	Offset int
}

//FetchOperands #
func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	bi.Offset = int(reader.ReadInt16())
}

//Index8Instruction #
type Index8Instruction struct {
	Index uint
}

//FetchOperands #
func (i8i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i8i.Index = uint(reader.ReadUint8())
}

//Index16Instruction #
type Index16Instruction struct {
	Index uint
}

//FetchOperands #
func (i16i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i16i.Index = uint(reader.ReadUint16())
}
