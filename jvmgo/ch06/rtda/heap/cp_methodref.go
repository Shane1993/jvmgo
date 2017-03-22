package heap

import "jvmgo/ch06/classfile"

type MethodRef struct {
	MemberRef
	method *Method 
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp 
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref//todo 此时class还没有被赋值，即还没被解析
}