package conversions

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type D2I struct { base.NoOperandsInstruction }
type D2L struct { base.NoOperandsInstruction }
type D2F struct { base.NoOperandsInstruction }

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushInt(int32(val))
}
func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushLong(int64(val))
}
func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushFloat(float32(val))
}