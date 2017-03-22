package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag	bool
	versionFlag	bool
	cpOption	string
	class		string
	args		[]string
}

func parseCmd() *Cmd {
	//创建一个新对象
	cmd := &Cmd{}

	flag.Usage = printUsage
	//设置需要解析的选项
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	//解析选项
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	//提示输入的格式
	fmt.Printf("Usage:.%s [-options] class [args...]n", os.Args[0])
}
