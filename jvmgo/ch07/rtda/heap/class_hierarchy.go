package heap

func (self *Class) IsAssignableFrom(cls *Class) bool {
	return self == cls ||
		self.IsSuperClassOf(cls) ||
		self.IsSuperInterfaceOf(cls)
}

// self implements iface
func (self *Class) IsImplements(iface *Class) bool {
	for k := self; k != nil; k = k.superClass {
		for _, i := range k.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// iface extends self
func (self *Class) IsSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}

// self extends iface
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends self
func (self *Class) IsSuperClassOf(c *Class) bool {
	return c.IsSubClassOf(self)
}

// self extends c
func (self *Class) IsSubClassOf(c *Class) bool {
	for k := self.superClass; k != nil; k = k.superClass {
		if k == c {
			return true
		}
	}
	return false
}
