package command

import "testing"

/*
1. 把每个命令执行以及需要的参数等东西都打包在一起，没有其他依赖，只需要触发即可。
2. 需要一个中间对象来触发（类似于一个面板）
3. 面板会提供一个统计的执行接口，方便统一处理（也就是实现的 command 接口）
*/

func TestCommand(t *testing.T) {
	electricCooker := new(ElectricCooker)

	electricCookerInvoker := new(ElectricCookerInvoker)

	steamRiceCommand := newSteamRiceCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(steamRiceCommand)
	println(electricCookerInvoker.ExecuteCookCommand())

	cookCongeeCommand := newCookCongeeCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(cookCongeeCommand)
	println(electricCookerInvoker.ExecuteCookCommand())

	shutdownCommand := newShutdownCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(shutdownCommand)
	println(electricCookerInvoker.ExecuteCookCommand())

}
