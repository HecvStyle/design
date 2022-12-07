package command

type CookCommand interface {
	Execute() string
}

type steamRiceCommand struct {
	electricCooker *ElectricCooker
}

func newSteamRiceCommand(electricCooker *ElectricCooker) *steamRiceCommand {
	return &steamRiceCommand{electricCooker: electricCooker}
}

func (s *steamRiceCommand) Execute() string {
	s.electricCooker.SetFire("中")
	s.electricCooker.setPressure("正常")
	return "蒸饭：" + s.electricCooker.Run("30分钟")
}

type cookCongeeCommand struct {
	electricCooker *ElectricCooker
}

func newCookCongeeCommand(electricCooker *ElectricCooker) *cookCongeeCommand {
	return &cookCongeeCommand{electricCooker: electricCooker}
}

func (c cookCongeeCommand) Execute() string {
	c.electricCooker.SetFire("大")
	c.electricCooker.setPressure("强")
	return "煮粥：" + c.electricCooker.Run("45分钟")
}

type shutdownCommand struct {
	electricCooker *ElectricCooker
}

func newShutdownCommand(electricCooker *ElectricCooker) *shutdownCommand {
	return &shutdownCommand{electricCooker: electricCooker}
}

func (s shutdownCommand) Execute() string {
	return s.electricCooker.ShutDown()
}

type ElectricCookerInvoker struct {
	cookCommand CookCommand
}

func (e *ElectricCookerInvoker) SetCookCommand(cookCommand CookCommand) {
	e.cookCommand = cookCommand
}

func (e *ElectricCookerInvoker) ExecuteCookCommand() string {

	return e.cookCommand.Execute()
}
