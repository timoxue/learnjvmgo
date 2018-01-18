package rtda

import "math"

//OperandStack #
type OperandStack struct {
	size  uint64
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

//PushInt #
func (operandStack *OperandStack) PushInt(val int32) {
	operandStack.slots[operandStack.size].num = val
	operandStack.size++
}

//PopInt #
func (operandStack *OperandStack) PopInt() int32 {
	operandStack.size--
	return operandStack.slots[operandStack.size].num
}

//PushFloat #
func (operandStack *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	operandStack.slots[operandStack.size].num = int32(bits)
	operandStack.size++
}

//PopFloat #
func (operandStack *OperandStack) PopFloat() float32 {
	operandStack.size--
	bits := uint32(operandStack.slots[operandStack.size].num)
	return math.Float32frombits(bits)
}

//PushLong #
func (operandStack *OperandStack) PushLong(val int64) {
	operandStack.slots[operandStack.size].num = int32(val)
	operandStack.slots[operandStack.size+1].num = int32(val >> 32)
	operandStack.size += 2
}

//PopLong #
func (operandStack *OperandStack) PopLong() int64 {
	operandStack.size -= 2
	low := uint32(operandStack.slots[operandStack.size].num)
	high := uint32(operandStack.slots[operandStack.size+1].num)
	return int64(high)<<32 | int64(low)
}

//PushDouble #
func (operandStack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	operandStack.PushLong(int64(bits))
}

//PopDouble #
func (operandStack *OperandStack) PopDouble() float64 {
	bits := uint64(operandStack.PopLong())
	return math.Float64frombits(bits)
}

//PushRef #
func (operandStack *OperandStack) PushRef(ref *Object) {
	operandStack.slots[operandStack.size].ref = ref
	operandStack.size++
}

//PopRef #
func (operandStack *OperandStack) PopRef() *Object {
	operandStack.size--
	ref := operandStack.slots[operandStack.size].ref
	operandStack.slots[operandStack.size].ref = nil
	return ref
}

//PushSlot #
func (operandStack *OperandStack) PushSlot(slot Slot) {
	operandStack.slots[operandStack.size] = slot
	operandStack.size++
}

//PopSlot #
func (operandStack *OperandStack) PopSlot() Slot {
	operandStack.size--
	return operandStack.slots[operandStack.size]
}
