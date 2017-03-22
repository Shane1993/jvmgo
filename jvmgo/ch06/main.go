package main

import "fmt"
import "strings"
import "jvmgo/ch06/classpath"
import "jvmgo/ch06/rtda/heap"

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

/**
 * 本节将使用ClassLoader来加载类
 * 测试相关的类与对象的指令
 */
func startJVM(cmd *Cmd) {

	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp)

	className := strings.Replace(cmd.class, ".", ".", -1)
	// cf :=loadClass(className, cp)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()//只取出main方法
	if mainMethod != nil {
		interpret(mainMethod)
	}else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}