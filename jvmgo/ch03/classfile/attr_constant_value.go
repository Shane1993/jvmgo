package classfile

/*
	该属性只会出现在field_info结构中
	ConstantValue_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 constantvalue_index;
	}
 */

type ConstantValueAttribute struct {
	constantvalueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantvalueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantvalueIndex() uint16 {
	return self.constantvalueIndex
}