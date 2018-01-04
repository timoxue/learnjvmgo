package rtda

//Frame #
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

//NewFrame #
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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
