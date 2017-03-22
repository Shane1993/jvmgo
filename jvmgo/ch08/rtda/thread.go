package rtda

import "jvmgo/ch08/rtda/heap"

type Thread struct {
	pc    int
	stack *Stack
}

/**
 * 创建Thread对象，设置栈帧的容量为1024
 */
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int      { return self.pc }
func (self *Thread) SetPC(pc int) { self.pc = pc }
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}
