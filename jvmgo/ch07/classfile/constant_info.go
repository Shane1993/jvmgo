package classfile

//定义14种常量tag
const (
	CONSTANT_Class				= 7
	CONSTANT_Fieldref			= 9
	CONSTANT_Methodref			= 10
	CONSTANT_InterfaceMethodref	= 11
	CONSTANT_String				= 8
	CONSTANT_Integer			= 3
	CONSTANT_Float				= 4
	CONSTANT_Long				= 5
	CONSTANT_Double				= 6
	CONSTANT_NameAndType		= 12
	CONSTANT_Utf8				= 1
	CONSTANT_MethodType			= 16
	CONSTANT_MethodHandle		= 15
	CONSTANT_InvokeDynamic		= 18
)


/*
	JVM中给出的常量结构
	cp_info {
		u1 tag;
		u1 info[];
	}
 */

/**
 * ConstantInfo接口中只有一个方法，就是解析方法读取常量信息
 * 	结构体中用于存储信息的具体字段由各自接口的具体实现结构体决定
 * 	所以接口中没有字段
 */
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	constantInfo := newConstantInfo(tag, cp)	//创建ConstantInfo
	constantInfo.readInfo(reader)	//初始化
	return constantInfo
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
		case CONSTANT_Integer				: return &ConstantIntegerInfo{}
		case CONSTANT_Float					: return &ConstantFloatInfo{}
		case CONSTANT_Long					: return &ConstantLongInfo{}
		case CONSTANT_Double				: return &ConstantDoubleInfo{}
		case CONSTANT_Utf8					: return &ConstantUtf8Info{}
		case CONSTANT_String				: return &ConstantStringInfo{cp : cp}
		case CONSTANT_Class 				: return &ConstantClassInfo{cp : cp}
		case CONSTANT_NameAndType			: return &ConstantNameAndTypeInfo{cp : cp}
		case CONSTANT_Fieldref				:
			return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp : cp}}
		case CONSTANT_Methodref 			:
			return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp : cp}}
		case CONSTANT_InterfaceMethodref	:
			return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp : cp}}
		case CONSTANT_MethodType			: return &ConstantMethodTypeInfo{}
		case CONSTANT_MethodHandle			: return &ConstantMethodHandleInfo{}
		case CONSTANT_InvokeDynamic			: return &ConstantInvokeDynamicInfo{}
		default : panic("java.lang.ClassFormatError: constant pool tag!")
	}
}

