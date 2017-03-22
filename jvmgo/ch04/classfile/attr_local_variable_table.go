package classfile 

/*
	LocalVariableTable_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 local_variable_table_length;
		{
			u2 start_pc;
			u2 line_number;
			u2 name_index;
			u2 descriptor_index;
			u2 index;
		} local_variable_table[local_variable_table_length];
	}
 */

type LocalVariableTableAttribute struct {
	localVariableTables []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc uint16
	length uint16
	nameIndex uint16
	descriptorIndex uint16 
	index uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUint16()
	self.localVariableTables = make([]*LocalVariableTableEntry, length)
	for i := range self.localVariableTables {
		self.localVariableTables[i] = &LocalVariableTableEntry {
			startPc : reader.readUint16(),
			length : reader.readUint16(),
			nameIndex : reader.readUint16(),
			descriptorIndex : reader.readUint16(),
			index : reader.readUint16(),
		}
	}
}
