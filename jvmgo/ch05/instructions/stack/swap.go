package stack

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val2)
}