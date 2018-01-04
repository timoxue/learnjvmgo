package classfile

//ConstantPool contains the constant information of the class
type ConstantPool []ConstantInfo

//ConstantInfo #
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANTInteger:
		return &ConstantIntegerInfo{}
	case CONSTANTFloat:
		return &ConstantFloatInfo{}
	case CONSTANTLong:
		return &ConstantLongInfo{}
	case CONSTANTDouble:
		return &ConstantDoubleInfo{}
	case CONSTANTUtf8:
		return &ConstantUtf8Info{}
	case CONSTANTString:
		return &ConstantStringInfo{cp: cp}
	case CONSTANTClass:
		return &ConstantClassInfo{cp: cp}
	case CONSTANTFieldref:
		return &ConstantFieldInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANTMethodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANTInterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANTNameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANTMethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANTMethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANTInvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(ntInfo.nameIndex)
	_type := cp.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
