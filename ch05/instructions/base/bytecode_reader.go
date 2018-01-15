package base

//BytecodeReader #
type BytecodeReader struct {
	code []byte
	pc   int
}

//Reset #
func (bcr *BytecodeReader) Reset(code []byte, pc int) {
	bcr.code = code
	bcr.pc = pc
}

//ReadUint8 #
func (bcr *BytecodeReader) ReadUint8() uint8 {
	i := bcr.code[bcr.pc]
	bcr.pc++
	return i
}

//ReadInt8 #
func (bcr *BytecodeReader) ReadInt8() int8 {
	return int8(bcr.ReadUint8())
}

//ReadUint16 #
func (bcr *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(bcr.ReadUint8())
	byte2 := uint16(bcr.ReadUint8())
	return (byte1 << 8) | byte2
}

//ReadInt16 #
func (bcr *BytecodeReader) ReadInt16() int16 {
	return int16(bcr.ReadUint16())
}

func (bcr *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(bcr.ReadUint8())
	byte2 := int32(bcr.ReadUint8())
	byte3 := int32(bcr.ReadUint8())
	byte4 := int32(bcr.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}
