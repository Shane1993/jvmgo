package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type ISHL struct { base.NoOperandsInstruction }
type LSHL struct { base.NoOperandsInstruction }
type ISHR struct { base.NoOperandsInstruction }
type LSHR struct { base.NoOperandsInstruction }

//逻辑右移
type IUSHR struct { base.NoOperandsInstruction }
type LUSHR struct { base.NoOperandsInstruction }

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f	//只需要5位就能表示移动31位了
	stack.PushInt(v1 << s)
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f	//只需要6位就能表示移动63位了
	stack.PushLong(v1 << s)
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f	
	stack.PushInt(v1 >> s)
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f	
	stack.PushLong(v1 >> s)
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	//将v1转换成无符号就能实现逻辑右移
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	//将v1转换成无符号就能实现逻辑右移
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}