package heap

import "jvmgo/ch06/classfile"

/**
 * 继承于ClassMember
 */
type Field struct {
	ClassMember
	slotId uint //用于记录该字段在Class的staticVars和Object的fields中的位置
	constValueIndex uint //
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) ConstantValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}