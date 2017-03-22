package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

/**
 * 该方法只用于解析接口索引表，因此先获取接口的个数，再解析多个索引
 */
func (self *ClassReader) readUint16s() []uint16 {
	length := self.readUint16()
	indexes := make([]uint16, length)
	for i := range indexes {
		indexes[i] = self.readUint16()
	}

	return indexes
}

/**
 * 读取指定长度的数据
 */
func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes

}