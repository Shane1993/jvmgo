package math

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type IMUL struct{ base.NoOperandsInstruction }
type FMUL struct{ base.NoOperandsInstruction }
type LMUL struct{ base.NoOperandsInstruction }
type DMUL struct{ base.NoOperandsInstruction }

func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val1 * val2)
}
func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	stack.PushFloat(val1 * val2)
}
func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val1 * val2)
}
func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	stack.PushDouble(val1 * val2)
}
