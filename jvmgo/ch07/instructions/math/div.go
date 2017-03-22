package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type IDIV struct { base.NoOperandsInstruction }
type FDIV struct { base.NoOperandsInstruction }
type LDIV struct { base.NoOperandsInstruction }
type DDIV struct { base.NoOperandsInstruction }

func (self *IDIV) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()

	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	stack.PushInt(val1 / val2)
}
func (self *FDIV) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	stack.PushFloat(val1 / val2)
}
func (self *LDIV) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()

	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	stack.PushLong(val1 / val2)
}
func (self *DDIV) Execute (frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	stack.PushDouble(val1 / val2)
}