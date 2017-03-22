package classpath

import "os"
import "path/filepath"
// import "fmt"

/**
 * Classpath用于存储和设定类的搜寻路径
 * 	在JVM加载类时便是通过Classpath获取到路径的
 */
type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

/**
 * 通过传入的jreOption, cpOption构造类的搜寻路径
 */
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

/**
 * 设置BootClasspath 和 ExtClasspath
 */
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	//先获取到jre的目录
	jreDir := getJreDir(jreOption)

	// fmt.Printf("classpath.go:  jreDir:%s\n", jreDir)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	//fmt.Printf("classpath.go: bootClasspath:%s\n", self.bootClasspath.String());
	
	// jre/lib/ext/*
	jreLibExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreLibExtPath)
	//fmt.Printf("classpath.go: extClasspath:%s\n", self.extClasspath.String());
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	//查找子目录是否含有jre
	if exists("./jre") {
		return "./jre"
	}

	//用JAVA_HOME作为jre目录
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder!")

}

func exists(path string) bool{
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/**
 * 设置UserClasspath
 */
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}

	self.userClasspath = newEntry(cpOption)
}



/**
 * 加载类，通过类的名字来获取class文件的二进制流
 * 	类似双亲委派模型，先从bootClasspath开始加载
 * 	然后从extClasspath加载
 * 	最后才从userClasspath加载
 */
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	//要先加上.class后缀
	className = className + ".class"

	// fmt.Printf("classpath.go:  className:%s\n", className);

	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		// fmt.Printf("classpath.go: 1\n");
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		// fmt.Printf("classpath.go: 2\n");
		return data, entry, err
	}
	// fmt.Printf("classpath.go: 3\n");
	return self.userClasspath.readClass(className)

}

/**
 * 返回类的用户路径名称
 */
func (self *Classpath) String() string {
	return self.userClasspath.String()
}