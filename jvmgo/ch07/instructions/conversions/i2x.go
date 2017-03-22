package conversions

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type I2L struct { base.NoOperandsInstruction }
type I2F struct { base.NoOperandsInstruction }
type I2D struct { base.NoOperandsInstruction }

type I2B struct { base.NoOperandsInstruction }
type I2C struct { base.NoOperandsInstruction }
type I2S struct { base.NoOperandsInstruction }

func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushLong(int64(val))
}
func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushFloat(float32(val))
}
func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushDouble(float64(val))
}

func (self *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(int32(int8(val)))
}
func (self *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(int32(uint16(val)))
}
func (self *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(int32(int16(val)))
}
