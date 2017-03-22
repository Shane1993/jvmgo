package control

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

/**
 * 注意该结构体没有使用Instruction接口，因此需要自行实现Instruction的接口方法
 *
 * 用于case值连续的switch
 */
type TABLE_SWITCH struct {
	defaultOffset int32	//记录defaultOffset的地址
	low int32
	high int32
	jumpOffsets []int32	//用于记录case值
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()	//保证defaultOffset地址是4的倍数
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	//先从操作数栈中取出输入值，从这里的PopInt可以看出switch里面只能传int类型参数
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index - self.low])
	} else {
		offset = int(self.defaultOffset)
	}

	base.Branch(frame, offset)
}