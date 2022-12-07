package state

import (
	"fmt"
	"testing"
)

/*
1. 把状态抽象处理，(满电，半电，空电)
2. 每个状态都有对应的处理，把处理和状态绑定在一起（connect和disconnect）
3. 没有设备都会有状态，三选一
4. 设备会因为一些操作触发状态变更，目标状态为当前操作引发的下一个状态，比如满电状态继续充电就会停止充电，进行过载保护
而如果是空电状态，接入电源后，下一个目的状态应该是满电，（我觉得也可以是部分电，具体可能还得看业务）
5. 主要就是理解状态的变化和流转，因为不同状态下的操作，下一个状态变化是需要根据当前状态决定的(有点交叉组合的味道)

*/

func TestState(t *testing.T) {
	iphone11 := NewIphone("11 pro")
	fmt.Println(iphone11.BatteryState())
	fmt.Println(iphone11.ConnectPlug())
	fmt.Println(iphone11.ConnectPlug())

	fmt.Println(iphone11.DisConnectPlug())
	fmt.Println(iphone11.DisConnectPlug())
	fmt.Println(iphone11.DisConnectPlug())

	fmt.Println(iphone11.ConnectPlug())
	fmt.Println(iphone11.DisConnectPlug())

}
