package stores

import (
	"learnjvmgo/ch05/instructions/base"
	"learnjvmgo/ch05/rtda"
)

//LStore #
type LStore struct{ base.Index8Instruction }

//LStore0 #
type LStore0 struct{ base.NoOperandsInstruction }

//LStore1 #
type LStore1 struct{ base.NoOperandsInstruction }

//LStore2 #
type LStore2 struct{ base.NoOperandsInstruction }

//LStore3 #
type LStore3 struct{ base.NoOperandsInstruction }

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

//Execute #
func (lstore *LStore) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(lstore.Index))
}

//Execute #
func (lstore2 *LStore2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}
