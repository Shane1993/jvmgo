package main

import "fmt"
import "jvmgo/ch07/instructions"
import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/heap"

/**
 * interpreter是一个解释器，作用是将方法的字节码转换成指令并执行
 * 因此作用域只是一个方法的字节码
 *
 * 因为方法是运行在thread里的一个栈帧，因此解释时需要thread和frame
 */

func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)

}

/**
 * 循环执行方法字节码
 */
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}

	for {
		//循环取出栈顶的栈帧，因为invoke指令会新增栈帧
		frame := thread.CurrentFrame()

		//先取出PC才能知道要执行的地方
		pc := frame.NextPC()
		//刷新thread中的PC
		thread.SetPC(pc)

		//decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		inst.Execute(frame)

		//在方法栈的栈帧为空时，说明线程的工作完毕，退出
		if thread.IsStackEmpty() {
			break
		}
	}
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}
