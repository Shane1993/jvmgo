package heap

/**
 * SymRef作为各符号引用的父类
 */
type SymRef struct {
	cp *ConstantPool //指向保存该符号引用的cp
	className string //存放符号本身所属的类的全限定名
	class *Class
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class 
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}

