package main

import "fmt"
// import "strings"
// import "jvmgo/ch04/classpath"
// import "jvmgo/ch04/classfile"
import "jvmgo/ch04/rtda" 

/**
 * 测试类加载的过程，主要关注startJVM这个方法
 */
func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	}else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	}else {
		startJVM(cmd)
	}
	
}

func startJVM(cmd *Cmd) {

	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
	
}

/**
 * Test the LocalVars
 */
func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 299792)
	vars.SetLong(4, -299792)
	vars.SetFloat(6, 3.14)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)

	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

/**
 * Test the OperandStack
 */
func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(299792)
	ops.PushLong(-299792)
	ops.PushFloat(3.14)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)

	//注意弹出的顺序与入栈的顺序相反
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}