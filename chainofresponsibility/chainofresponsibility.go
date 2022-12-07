package chainofresponsibility

import "fmt"

// 定义登机过程中的每一个处理
type BoardingProcessor interface {
	// 处理完了，需要交给下一个处理流程
	SetNextProcessor(processor BoardingProcessor)
	// 处理的流程
	ProcessFor(passenger *Passenger)
}

// 旅客
type Passenger struct {
	name                  string
	hasBoardingPass       bool //是否办理登机牌
	hasLuggage            bool // 是否有托运心里
	isPassIdentityCheck   bool //是否能通过身份验证
	isPassSecurityCheck   bool // 是否通过安检
	isCompleteForBoarding bool //是否完成了登机
}

type baseBoardingProcessor struct {
	nextProcessor BoardingProcessor
}

func (b *baseBoardingProcessor) SetNextProcessor(processor BoardingProcessor) {
	b.nextProcessor = processor
}

func (b *baseBoardingProcessor) ProcessFor(passenger *Passenger) {
	if b.nextProcessor != nil {
		b.nextProcessor.ProcessFor(passenger)
	}
}

type boardingPassProcessor struct {
	baseBoardingProcessor
}

func (b *boardingPassProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("为旅客%s办理登机牌\n", passenger.name)
		passenger.hasBoardingPass = true
	}
	b.baseBoardingProcessor.ProcessFor(passenger)
}

type luggageCheckInProcessor struct {
	baseBoardingProcessor
}

func (l *luggageCheckInProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客%s未办理登机牌,无法托运行李;\n", passenger.name)
		return
	}
	if passenger.hasLuggage {
		fmt.Printf("为旅客%s办理行李托运;\n", passenger.name)
		passenger.hasLuggage = true
	}
	l.baseBoardingProcessor.ProcessFor(passenger)
}

type identifyCheckProcessor struct {
	baseBoardingProcessor
}

func (i *identifyCheckProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客%s未办理登机牌,无法办理身份检测;\n", passenger.name)
		return
	}
	if !passenger.isPassIdentityCheck {
		fmt.Printf("为旅客%s办理身份验证;\n", passenger.name)
		passenger.isPassIdentityCheck = true
	}
	i.baseBoardingProcessor.ProcessFor(passenger)
}

type securityCheckProcess struct {
	baseBoardingProcessor
}

func (s *securityCheckProcess) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客%s未办理登机牌,无法进行安检;\n", passenger.name)
		return
	}
	if !passenger.isPassSecurityCheck {
		fmt.Printf("为旅客%s办理安检;\n", passenger.name)
		passenger.isPassSecurityCheck = true
	}
	s.baseBoardingProcessor.ProcessFor(passenger)
}

type completeBoardingProcessor struct {
	baseBoardingProcessor
}

func (c *completeBoardingProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass ||
		!passenger.isPassIdentityCheck ||
		!passenger.isPassSecurityCheck {
		fmt.Printf("旅客%s登机检测未完成，不能登机;\n", passenger.name)
		return
	}
	passenger.isCompleteForBoarding = true
	fmt.Printf("旅客%s登机检测完成，成功登机;\n", passenger.name)
}
