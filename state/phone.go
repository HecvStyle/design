package state

import "fmt"

var (
	FullBatteryState  = new(fullBatteryState)  // 满电状态
	PartBatteryState  = new(partBatteryState)  // 部分电量状态
	EmptyBatteryState = new(emptyBatteryState) // 空电状态
)

type Iphone struct {
	model        string
	batteryState BatteryState
}

func NewIphone(model string) *Iphone {
	return &Iphone{model: model, batteryState: PartBatteryState}
}

func (i *Iphone) BatteryState() string {
	return fmt.Sprintf("iphone: %s 当前电量为：%s", i.model, i.batteryState)
}

func (i *Iphone) ConnectPlug() string {
	return fmt.Sprintf("iphone %s连接电源,%s", i.model, i.batteryState.ConnectPlug(i))
}

func (i *Iphone) DisConnectPlug() string {
	return fmt.Sprintf("iphone %s断开电源,%s", i.model, i.batteryState.DisConnectPlug(i))
}

func (i *Iphone) SetBatteryState(state BatteryState) {
	i.batteryState = state
}

func (i *Iphone) charge() string {
	return "正在充电"
}

func (i *Iphone) pauseCharge() string {

	return "停止充电"
}

func (i *Iphone) shutdown() string {
	return "没电关机"
}

func (i *Iphone) consume() string {
	return "使用中，持续消耗电量"
}
