package classfile

//MemberInfo #
type MemberInfo struct {
	cp              ConstantPool
	accessFlag      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlag:      reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

//Name #
func (memberInfo *MemberInfo) Name() string {
	return memberInfo.cp.getUtf8(memberInfo.nameIndex)
}

//Descriptor #
func (memberInfo *MemberInfo) Descriptor() string {
	return memberInfo.cp.getUtf8(memberInfo.descriptorIndex)
}
