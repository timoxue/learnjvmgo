package classfile

//SourceFileAttribute #
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (sfa *SourceFileAttribute) readInfo(reader *ClassReader) {
	sfa.sourceFileIndex = reader.readUint16()
}

//FileName #
func (sfa *SourceFileAttribute) FileName() string {
	return sfa.cp.getUtf8(sfa.sourceFileIndex)
}
