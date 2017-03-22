package base

import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/heap"

func InitClass(thread *rtda.Thread, class *heap.Class) {
	//设置标志位
	class.StartInit()
	//将类初始化方法压入新栈帧
	scheduleClinit(thread, class)
	//初始化父类
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
