package classfile

/*
	JVM中CONSTANT_NameAndType_info结构
	CONSTANT_NameAndType_info {
		u1 tag;
		u2 nameIndex;
		u2 typeIndex;
	}
 */

type ConstantNameAndTypeInfo struct {
	cp ConstantPool
	nameIndex uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
