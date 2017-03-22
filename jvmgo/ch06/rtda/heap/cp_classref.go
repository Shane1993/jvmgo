package heap

import "jvmgo/ch06/classfile"

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp 
	ref.className = classInfo.Name()
	return ref //todo 此时class还没有被赋值，即还没被解析
}


