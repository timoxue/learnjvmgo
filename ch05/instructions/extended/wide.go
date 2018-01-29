package extended

import (
	"jvmgo-book/v1/code/go/src/jvmgo/ch05/instructions/loads"
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//Wide #
type Wide struct {
	modifiedInstruction base.Instruction
}

//FetchOperands #
func (wide *Wide) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		reader.modifiedInstruction = inst
	case 0x16:
	case 0x17:
	case 0x18:
	case 0x19:
	case 0x36:
	case 0x37:
	case 0x38:
	case 0x39:
	case 0x3a:
	case 0x84:
		inst := &match.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		wide.modifiedInstuction = inst
	case 0xa9:
		panic("Unsupported opcode: 0xa9!")
	}
}

//Execute #
func (wide *Wide) Execute(frame *rtda.Frame) {
	wide.modifiedInstruction.Execute(frame)
}
