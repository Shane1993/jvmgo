package math

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type ISUB struct{ base.NoOperandsInstruction }
type FSUB struct{ base.NoOperandsInstruction }
type LSUB struct{ base.NoOperandsInstruction }
type DSUB struct{ base.NoOperandsInstruction }

func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val2 - val1)
}
func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	stack.PushFloat(val2 - val1)
}
func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val2 - val1)
}
func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	stack.PushDouble(val2 - val1)
}
