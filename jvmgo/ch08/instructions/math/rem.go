package math

import "math"
import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type IREM struct{ base.NoOperandsInstruction }
type FREM struct{ base.NoOperandsInstruction }
type LREM struct{ base.NoOperandsInstruction }
type DREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()

	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	stack.PushInt(val1 % val2)
}
func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	stack.PushFloat(float32(math.Mod(float64(val1), float64(val2))))
}
func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()

	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	stack.PushLong(val1 % val2)
}
func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	stack.PushDouble(math.Mod(val1, val2))
}
