package classfile

//ConstantPool实际上是由ConstantInfo数组构成的，因此ConstantPool并没有实际结构
type ConstantPool []ConstantInfo


/**
 * 用于返回一个解析完成的ConstantPool
 */
func readConstantPool(reader *ClassReader) ConstantPool {
	count := int(reader.readUint16())
	cp := make([]ConstantInfo, count)

	//注意常量池的常量索引由1开始
	for i := 1; i < count; i++ {
		cp[i] = readConstantInfo(reader, cp)
		//long和double两种类型占据两个位置，因此遇到时索引值要加1
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}

	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}