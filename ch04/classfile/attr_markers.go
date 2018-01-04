package classfile

//DeprecatedAttribute #
type DeprecatedAttribute struct {
	MarkerAttribute
}

//SyntheticAttribute #
type SyntheticAttribute struct {
	MarkerAttribute
}

//MarkerAttribute #
type MarkerAttribute struct {
}

func (ma *MarkerAttribute) readInfo(reader *ClassReader) {
	//read nothing
}
