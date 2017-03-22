package math

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"

//iinc指令是给局部变量表中的int变量增加常量值，因此需要局部变量表的索引和常量值作为字段
type IINC struct {
	Index uint
	Const int32
}

/**
 * 由于IINC并没有直接利用Instruction的抽象父类，因此需要自己实现Instruction的接口方法
 */
func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	vars := frame.LocalVars()
	val := vars.GetInt(self.Index)
	val += self.Const
	vars.SetInt(self.Index, val)
}
