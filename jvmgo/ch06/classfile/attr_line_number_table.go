package classfile 

/*
	LineNumberTable_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 line_number_table_length;
		{
			u2 start_pc;
			u2 line_number;
		} line_number_table[line_number_table_length];
	}
 */

type LineNumberTableAttribute struct {
	lineNumberTables []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUint16()
	self.lineNumberTables = make([]*LineNumberTableEntry, length)
	for i := range self.lineNumberTables {
		self.lineNumberTables[i] = &LineNumberTableEntry {
			startPc : reader.readUint16(),
			lineNumber : reader.readUint16(),
		}
	}
}
