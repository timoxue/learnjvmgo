package classfile

//ConstantPool contains the constant information of the class
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantPool, cpCount)

}
