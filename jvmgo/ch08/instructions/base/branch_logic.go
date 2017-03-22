package base

import "jvmgo/ch08/rtda"

/**
 * 该方法是跳转逻辑，在跳转指令中会用到，用于设置下次frame要执行的pc
 */
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	pc += offset
	frame.SetNextPC(pc)
}
