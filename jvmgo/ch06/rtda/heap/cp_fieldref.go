package heap

import "jvmgo/ch06/classfile"

type FieldRef struct {
	MemberRef
	field *Field //直接引用
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref//todo 此时class还没有被赋值，即还没被解析
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class 
	c := self.ResolvedClass()
	field := lookupField(c, self.name, self.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	//验证权限
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

/**
 * 搜寻字段
 * Class是搜寻的类名，找的特征是name和descriptor
 */
func lookupField(c *Class, name, descriptor string) *Field {


	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	return nil
}