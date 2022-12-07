package mediator

type AirportMediator interface {
	CanLandAirport(aircraft AirCraft) bool // 当前飞机是否可以将来
	NotifyWaitingAircraft()                // 通知等待的飞机
}

type AirportTower struct {
	hasFreeAirstrip bool
	waitingQueue    []AirCraft
}

func (a *AirportTower) CanLandAirport(aircraft AirCraft) bool {
	if a.hasFreeAirstrip {
		a.hasFreeAirstrip = false
		return true
	}

	// 进入等待队列，获取通知后可以降落
	a.waitingQueue = append(a.waitingQueue, aircraft)
	return false
}

func (a *AirportTower) NotifyWaitingAircraft() {
	if !a.hasFreeAirstrip {
		a.hasFreeAirstrip = true
	}
	if len(a.waitingQueue) > 0 {
		first := a.waitingQueue[0]
		a.waitingQueue = a.waitingQueue[1:]
		first.ApproachAirport()
	}
}
