package rtda

//Frame #
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
}

//NewFrame #
func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

//LocalVars #
func (frame *Frame) LocalVars() LocalVars {
	return frame.localVars
}

//OperandStack #
func (frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}

//Thread #
func (frame *Frame) Thread() *Thread {
	return frame.thread
}

//NextPC #
func (frame *Frame) NextPC() int {
	return frame.nextPC
}

//SetNextPC #
func (frame *Frame) SetNextPC(nextPC int) {
	frame.nextPC = nextPC
}
