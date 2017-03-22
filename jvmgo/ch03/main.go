package main

import "fmt"
import "strings"
import "jvmgo/ch03/classpath"
import "jvmgo/ch03/classfile"

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

	// fmt.Printf("XjreOption:" + cmd.XjreOption + ", cpOption:" + cmd.cpOption + "\n")
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	
	fmt.Printf(cmd.class)
	printClassInfo(cf)
	
}

/**
 * 通过className在搜寻路径classpath中获取到class文件
 * 	并将其解析到ClassFile结构体当中返回
 */
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

/**
 * 通过ClassFile结构体输出类的相关信息
 */
func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n",cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, field := range cf.Fields() {
		fmt.Printf("  %s\n", field.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, method := range cf.Methods() {
		fmt.Printf("  %s\n", method.Name())
	}
}