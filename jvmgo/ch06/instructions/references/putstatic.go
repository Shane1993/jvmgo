package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

/**
 * 需要两个操作数，一个是某个静态变量的索引
 * 	另一个是操作数栈顶的数值
 */
type PUT_STATIC struct {
	base.Index16Instruction
}

/**
 * 1 先通过获取FieldRef（首先还要获取到ConstantPool）
 * 2 解析成Field
 * 3 获取该Field的slotId
 * 4 从操作数栈中取出数值存入staticVars中与slotId对应的位置
 */
func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I': slots.SetInt(slotId, stack.PopInt())
	case 'F': slots.SetFloat(slotId, stack.PopFloat())
	case 'J': slots.SetLong(slotId, stack.PopLong())
	case 'D': slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[': slots.SetRef(slotId, stack.PopRef())
	}
}

