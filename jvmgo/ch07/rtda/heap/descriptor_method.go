package heap

/**
 * 用于记录方法中多个形参以及返回值的类型
 */
type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (self *MethodDescriptor) ParameterTypes() []string {
	return self.parameterTypes
}
func (self *MethodDescriptor) ReturnType() string {
	return self.returnType
}

func (self *MethodDescriptor) addParameterType(t string) {
	pLen := len(self.parameterTypes)
	if pLen == cap(self.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, self.parameterTypes)
		self.parameterTypes = s
	}

	self.parameterTypes = append(self.parameterTypes, t)
}
