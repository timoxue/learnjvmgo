package classfile

//ConstantMethodTypeInfo #
type ConstantMethodTypeInfo struct{}

//ConstantMethodHandleInfo #
type ConstantMethodHandleInfo struct{}

//ConstantInvokeDynamicInfo #
type ConstantInvokeDynamicInfo struct{}

func (cmtf *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {}

func (cmhf *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {}

func (cidi *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {}
