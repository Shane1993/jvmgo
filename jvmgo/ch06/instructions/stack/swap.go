package stack

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

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