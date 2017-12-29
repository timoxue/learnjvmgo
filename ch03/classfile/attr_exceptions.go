package classfile

//ExceptionsAttribute #
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (ea *ExceptionsAttribute) readInfo(reader *ClassReader) {
	ea.exceptionIndexTable = reader.readUint16s()
}

//ExceptionIndexTable #
func (ea *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return ea.exceptionIndexTable
}
