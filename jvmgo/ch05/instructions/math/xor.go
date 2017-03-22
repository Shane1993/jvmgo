package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type IXOR struct { base.NoOperandsInstruction }
type LXOR struct { base.NoOperandsInstruction }

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v1 ^ v2)
}
func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	stack.PushLong(v1 ^ v2)
}