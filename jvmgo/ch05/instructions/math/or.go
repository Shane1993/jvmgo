package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type IOR struct { base.NoOperandsInstruction }
type LOR struct { base.NoOperandsInstruction }

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v1 | v2)
}
func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	stack.PushLong(v1 | v2)
}