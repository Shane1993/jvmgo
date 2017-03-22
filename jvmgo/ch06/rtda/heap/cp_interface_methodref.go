package heap

import "jvmgo/ch06/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method 
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp 
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref//todo 此时class还没有被赋值，即还没被解析
}