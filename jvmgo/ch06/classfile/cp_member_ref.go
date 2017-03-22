package classfile

/*
	JVM中字段Field_info的结构，Method，InterfaceMethod一样
	CONSTANT_Field_info {
		u1 tag;
		u2 classIndex;
		u2 nameAndTypeIndex;
	}
 */

type ConstantMemberrefInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct { ConstantMemberrefInfo }
type ConstantMethodrefInfo struct { ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct { ConstantMemberrefInfo }