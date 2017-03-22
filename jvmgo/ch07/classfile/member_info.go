package classfile

type MemberInfo struct {
	cp					ConstantPool
	accessFlags			uint16
	nameIndex			uint16
	descriptorIndex		uint16
	attributes			[]AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	count := reader.readUint16()
	members := make([]*MemberInfo, count)
	for i := range members {
		members[i] = readMember(reader, cp)
	}

	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp 				:	cp,
		accessFlags 	:	reader.readUint16(),
		nameIndex 		: 	reader.readUint16(),
		descriptorIndex : 	reader.readUint16(),
		attributes 		: 	readAttributes(reader, cp),
	}
}

/**
 * 返回成员的访问标志位
 */
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

/**
 * 从多个属性中找出方法属性，方法的字节码存在与方法属性中
 */
func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}
/**
 * 从多个属性中找出ConstantValue属性，final属于该属性
 */
func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

