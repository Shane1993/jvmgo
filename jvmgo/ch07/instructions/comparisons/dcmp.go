package comparisons

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type DCMPL struct { base.NoOperandsInstruction }
type DCMPG struct { base.NoOperandsInstruction }

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 < v2 {
		stack.PushInt(-1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 > v2{
		stack.PushInt(1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}
func (self *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}