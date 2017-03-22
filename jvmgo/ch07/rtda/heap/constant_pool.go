package heap

import "fmt"
import "jvmgo/ch07/classfile"

//用于接收运行时常量池的常量
type Constant interface {}

type ConstantPool struct {
	class *Class 
	consts []Constant
}

/**
 * 将文件常量池cfCp转化为运行时常量池rtCp
 */
func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}

	//注意常量池中的索引从1开始
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			consts[i] = cpInfo.(*classfile.ConstantIntegerInfo).Value()
		case *classfile.ConstantFloatInfo:
			consts[i] = cpInfo.(*classfile.ConstantFloatInfo).Value()
		case *classfile.ConstantLongInfo:
			consts[i] = cpInfo.(*classfile.ConstantLongInfo).Value()
			i++
		case *classfile.ConstantDoubleInfo:
			consts[i] = cpInfo.(*classfile.ConstantDoubleInfo).Value()
			i++
		case *classfile.ConstantStringInfo:
			consts[i] = cpInfo.(*classfile.ConstantStringInfo).String()
		case *classfile.ConstantUtf8Info:
			consts[i] = newConstantUtf8(cpInfo.(*classfile.ConstantUtf8Info))
		case *classfile.ConstantClassInfo:
			consts[i] = newClassRef(rtCp, cpInfo.(*classfile.ConstantClassInfo))
		case *classfile.ConstantFieldrefInfo:
			consts[i] = newFieldRef(rtCp, cpInfo.(*classfile.ConstantFieldrefInfo))
		case *classfile.ConstantMethodrefInfo:
			consts[i] = newMethodRef(rtCp, cpInfo.(*classfile.ConstantMethodrefInfo))
		case *classfile.ConstantInterfaceMethodrefInfo:
			consts[i] = newInterfaceMethodRef(rtCp, cpInfo.(*classfile.ConstantInterfaceMethodrefInfo))
		// case *classfile.ConstantInvokeDynamicInfo:
		// 	consts[i] = newConstantInvokeDynamic(rtCp, cpInfo.(*classfile.ConstantInvokeDynamicInfo))
		// case *classfile.ConstantMethodHandleInfo:
		// 	consts[i] = newConstantMethodHandle(cpInfo.(*classfile.ConstantMethodHandleInfo))
		// case *classfile.ConstantMethodTypeInfo:
		// 	consts[i] = newConstantMethodType(cpInfo.(*classfile.ConstantMethodTypeInfo))
		default:
			// todo
			//fmt.Printf("%T \n", cpInfo)
			// panic("todo")			
		}
	}
	return rtCp
}

/**
 * 根据index提供常量池中的常量
 */
func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constant at index %d", index))
}

