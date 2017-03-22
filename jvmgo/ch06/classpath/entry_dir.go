package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	//包的绝对路径，注，这里并没有包含实际类名
	absDir string
}

//创建一个DirEntry对象
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	//通过构造体创建DirEntry对象
	return &DirEntry{absDir}
}

//实现读取类信息的接口方法
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

//实现自身信息的接口方法
func (self *DirEntry) String() string {
	return self.absDir
}
