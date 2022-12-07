package mediator

import "testing"

/*
这里就是个线性的操作了，保证了FIFO
所有的调协都有mediator来做，这样来看就集中了，但如果抽象不好，那mediator的逻辑会变得非常复杂
*/
func TestMediator(t *testing.T) {
	airportMediator := &AirportTower{
		hasFreeAirstrip: true,
		waitingQueue:    nil,
	}

	c919 := newAirliner("C-919", airportMediator)

	m26 := newHelicopter("M-26", airportMediator)

	c919.ApproachAirport()
	m26.ApproachAirport()

	c919.DepartAirport()
	m26.DepartAirport()

}
