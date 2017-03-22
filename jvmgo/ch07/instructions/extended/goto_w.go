package extended

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

type GOTO_W struct {
	Offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.Offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}