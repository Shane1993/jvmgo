package base

import "jvmgo/ch07/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}


/**
 * NoOperandsInstruction是所有无操作数指令的抽象“父类”
 */
type NoOperandsInstruction struct {

}
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	//do nothing
}

/**
 * 跳转指令
 */
type BranchInstruction struct{
	Offset int
}
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/**
 * 局部变量表索引指令，大部分索引以uint8形式存在字节码中
 */
type Index8Instruction struct {
	Index uint
}
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

/**
 * 操作数栈索引指令，uint16
 */
type Index16Instruction struct {
	Index uint
}
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}





