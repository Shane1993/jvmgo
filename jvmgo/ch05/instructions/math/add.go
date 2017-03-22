package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type IADD struct { base.NoOperandsInstruction }
type FADD struct { base.NoOperandsInstruction }
type LADD struct { base.NoOperandsInstruction }
type DADD struct { base.NoOperandsInstruction }

func (self *IADD) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val1 + val2)
}
func (self *FADD) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	stack.PushFloat(val1 + val2)
}
func (self *LADD) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val1 + val2)
}
func (self *DADD) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	stack.PushDouble(val1 + val2)
}