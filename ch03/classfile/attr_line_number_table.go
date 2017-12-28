package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (lineNumberTableAttribute *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	lineNumberTableAttribute.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range lineNumberTableAttribute.lineNumberTable {
		lineNumberTableAttribute.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
