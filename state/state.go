package state

type BatteryState interface {
	ConnectPlug(iphone *Iphone) string
	DisConnectPlug(iphone *Iphone) string
}

type fullBatteryState struct {
}

func (s *fullBatteryState) String() string {
	return "满电状态"
}

func (s *fullBatteryState) ConnectPlug(iphone *Iphone) string {
	iphone.pauseCharge()
	return "满电状态，停止充电"
}

func (s *fullBatteryState) DisConnectPlug(iphone *Iphone) string {
	iphone.consume()
	iphone.SetBatteryState(PartBatteryState)
	return "进入耗电状态"
}

type emptyBatteryState struct {
}

func (s *emptyBatteryState) String() string {
	return "空电状态"
}

func (s *emptyBatteryState) ConnectPlug(iphone *Iphone) string {
	// 进行充电，并转变为部分电量状态
	iphone.SetBatteryState(PartBatteryState)
	iphone.charge()
	return "正在进行充电"
}

// 空电状态下断开电源，那肯定是关机
func (s *emptyBatteryState) DisConnectPlug(iphone *Iphone) string {
	iphone.shutdown()
	return "关机"
}

type partBatteryState struct {
}

func (s *partBatteryState) String() string {
	return "部分电量状态"
}

func (s *partBatteryState) ConnectPlug(iphone *Iphone) string {
	// 进行充电，并转变为部分电量状态
	iphone.SetBatteryState(FullBatteryState)
	iphone.charge()
	return "正在进行充电,知道满电为止"
}

func (s *partBatteryState) DisConnectPlug(iphone *Iphone) string {

	iphone.SetBatteryState(EmptyBatteryState)
	iphone.consume()
	return "结束充电，进入耗电状态"
}
