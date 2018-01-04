package classfile

//ConstantStringInfo #
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (csi *ConstantStringInfo) readInfo(reader *ClassReader) {
	csi.stringIndex = reader.readUint16()
}

//String #
func (csi *ConstantStringInfo) String() string {
	return csi.cp.getUtf8(csi.stringIndex)
}
