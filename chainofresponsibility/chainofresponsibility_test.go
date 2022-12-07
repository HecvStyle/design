package chainofresponsibility

import "testing"

/*
所有的操作任务都被分割成为单个节点
各个节点通过指针做成了链表
每个节点的方法都是相同的，但处理的逻辑是不同的。因为节点抽象的，各自的具体实现在实现类里
需要触发头节点来遍历触发整个节点的处理方法（链表的遍历）
通过节点属性控制该节点是否需要被触发处理操作
*/

func TestChainOfResponsibility(t *testing.T) {
	// 构造责任链
	boardingProcessor := BuildBoardingProcessorChain()
	passenger := &Passenger{
		name:                  "Bob",
		hasBoardingPass:       false,
		hasLuggage:            false,
		isPassIdentityCheck:   false,
		isPassSecurityCheck:   false,
		isCompleteForBoarding: false,
	}
	// 通过头节点触发连续的调用
	//
	boardingProcessor.ProcessFor(passenger)

}

func BuildBoardingProcessorChain() BoardingProcessor {
	completeBoardingNode := &completeBoardingProcessor{}

	securityCheckNode := &securityCheckProcess{}
	securityCheckNode.SetNextProcessor(completeBoardingNode)

	identifyCheckNode := &identifyCheckProcessor{}
	identifyCheckNode.SetNextProcessor(securityCheckNode)

	luggageCheckinNode := &luggageCheckInProcessor{}
	luggageCheckinNode.SetNextProcessor(identifyCheckNode)

	boardingPassNode := &boardingPassProcessor{}
	boardingPassNode.SetNextProcessor(luggageCheckinNode)

	return boardingPassNode
}
