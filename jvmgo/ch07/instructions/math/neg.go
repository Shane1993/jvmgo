package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type INEG struct { base.NoOperandsInstruction }
type FNEG struct { base.NoOperandsInstruction }
type LNEG struct { base.NoOperandsInstruction }
type DNEG struct { base.NoOperandsInstruction }

func (self *INEG) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}
func (self *FNEG) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}
func (self *LNEG) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
func (self *DNEG) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}