package conversions

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

type F2I struct{ base.NoOperandsInstruction }
type F2L struct{ base.NoOperandsInstruction }
type F2D struct{ base.NoOperandsInstruction }

func (self *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushInt(int32(val))
}
func (self *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushLong(int64(val))
}
func (self *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushDouble(float64(val))
}
