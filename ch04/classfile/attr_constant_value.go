package classfile

//ConstantValueAttribute #
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (cva *ConstantValueAttribute) readInfo(reader *ClassReader) {
	cva.constantValueIndex = reader.readUint16()
}

//ConstantValueIndex #
func (cva *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return cva.constantValueIndex
}
