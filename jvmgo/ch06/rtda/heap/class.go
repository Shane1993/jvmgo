package heap

import "jvmgo/ch06/classfile"
import "strings"

/**
 * Class结构体，通过ClassFile解析出来
 */
type Class struct {
	AccessFlags
	name string //thisClassName
	superClassName string
	interfaceNames []string
	constantPool *ConstantPool //运行时常量池，注意并不是classfile中的常量池
	fields []*Field 
	methods []*Method 
	loader *ClassLoader
	superClass *Class //superClass是通过superClassName解析出来的Class
	interfaces []*Class 
	instanceSlotCount uint //实例字段个数
	staticSlotCount uint //静态字段个数
	staticVars Slots //静态字段表
}


func (self *Class) String() string {
	return "{Class name:" + self.name + "}"
}

// getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) Name() string {
	return self.name
}
func (self *Class) Methods() []*Method {
	return self.methods
}
func (self *Class) Fields() []*Field {
	return self.fields
}
func (self *Class) Loader() *ClassLoader {
	return self.loader
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) Interfaces() []*Class {
	return self.interfaces
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

/**
 * 通过ClassFile解析成Class
 */
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())//在ConstantPool中
	class.fields = newFields(class, cf.Fields())//在Filed中
	class.methods = newMethods(class, cf.Methods())//在Method中
	return class
}

func (self *Class) isAccessibleTo(other *Class) bool {
	return self.IsPublic() || self.getPackageName() == other.getPackageName()
}

func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, ""); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

/** 
 * 根据Class创建Object
 */
func newObject(class *Class) *Object {
	return &Object{
		class: class, 
		fields: newSlots(class.instanceSlotCount),
	}
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name && method.descriptor == descriptor {

				return method
			}
	}
	return nil
}