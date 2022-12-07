package command

import "fmt"

type ElectricCooker struct {
	fire     string // 火力
	pressure string // 压力
}

func (e *ElectricCooker) SetFire(fire string) {
	e.fire = fire
}

func (e *ElectricCooker) setPressure(pressure string) {
	e.pressure = pressure
}

func (e *ElectricCooker) Run(duration string) string {

	return fmt.Sprintf("电饭煲设置火力：%s,压力：%s,时间：%s", e.fire, e.pressure, duration)
}

func (e *ElectricCooker) ShutDown() string {

	return "电饭煲已停止运行"
}
