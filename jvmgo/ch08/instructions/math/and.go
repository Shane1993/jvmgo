package math

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type IAND struct{ base.NoOperandsInstruction }
type LAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v1 & v2)
}
func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	stack.PushLong(v1 & v2)
}
