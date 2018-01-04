package classfile

//ConstantMemberrefInfo #
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (cmi *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	cmi.classIndex = reader.readUint16()
	cmi.nameAndTypeIndex = reader.readUint16()
}

//ClassName #
func (cmi *ConstantMemberrefInfo) ClassName() string {
	return cmi.cp.getClassName(cmi.classIndex)
}

//NameAndDescriptor +
func (cmi *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return cmi.cp.getNameAndType(cmi.nameAndTypeIndex)
}

//ConstantFieldInfo +
type ConstantFieldInfo struct{ ConstantMemberrefInfo }

//ConstantMethodrefInfo +
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }

//ConstantInterfaceMethodrefInfo +
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }
