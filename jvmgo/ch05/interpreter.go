package main 

import "fmt"
import "jvmgo/ch05/classfile"
import "jvmgo/ch05/instructions"
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/**
 * interpreter是一个解释器，作用是将方法的字节码转换成指令并执行
 * 因此作用域只是一个方法的字节码
 *
 * 因为方法是运行在thread里的一个栈帧，因此解释时需要thread和frame
 */

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, bytecode)

}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n",frame.LocalVars())
		fmt.Printf("OperandStack:%v\n",frame.OperandStack())
		panic(r)
	}
}

/**
 * 循环执行方法字节码
 */
func loop(thread *rtda.Thread, bytecode []byte) {
	//先取出栈帧
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		//先取出PC才能知道要执行的地方
		pc := frame.NextPC()
		//刷新thread中的PC
		thread.SetPC(pc)

		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}