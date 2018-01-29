package rtda

//Thread #
type Thread struct {
	pc    int
	stack *Stack
}

//NewThread #
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

//PushFrame #
func (thread *Thread) PushFrame(frame *Frame) {
	thread.stack.push(frame)
}

//PopFrame #
func (thread *Thread) PopFrame() *Frame {
	return thread.stack.pop()
}

//CurrentFrame #
func (thread *Thread) CurrentFrame() *Frame {
	return thread.stack.top()
}

//NewFrame #
func (thread *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(thread, maxLocals, maxStack)
}

//PC #
func (thread *Thread) PC() int {
	return thread.pc
}

//SetPC #
func (thread *Thread) SetPC(pc int) {
	thread.pc = pc
}
