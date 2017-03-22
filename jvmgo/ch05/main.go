package main

import "fmt"
import "strings"
import "jvmgo/ch05/classpath"
import "jvmgo/ch05/classfile"

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

	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", ".", -1)
	cf :=loadClass(className, cp)
	mainMethod := getMainMethod(cf)//只取出main方法
	if mainMethod != nil {
		interpret(mainMethod)
	}else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

	
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}

	return nil
}