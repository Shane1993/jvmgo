package classfile

// import "fmt"
// import "unicode/utf16"

/*
	JVM中Utf8_info的结构体
	CONSTANT_Utf8_info {
		u1 tag;
		u2 length;
		u1 byte[length];
	}
 */

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func (self *ConstantUtf8Info) Str() string {
	return self.str
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}

