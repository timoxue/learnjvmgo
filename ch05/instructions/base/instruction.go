package base

import "learnjvmgo/ch05/rtda"

type Instruction interface {
	FetchOprands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}
