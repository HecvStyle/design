package mediator

import "fmt"

type AirCraft interface {
	ApproachAirport() // 达到机场
	DepartAirport()   // 飞离机场
}

type airliner struct {
	name            string
	airportMediator AirportMediator
}

func newAirliner(name string, airportMediator AirportMediator) *airliner {
	return &airliner{name: name, airportMediator: airportMediator}
}

func (a *airliner) ApproachAirport() {
	if !a.airportMediator.CanLandAirport(a) {
		fmt.Printf("机场繁忙，%s请继续等待\n", a.name)
		return
	}
	fmt.Printf("%s已降落\n", a.name)

}

func (a *airliner) DepartAirport() {
	fmt.Printf("%s已经起飞离开机场;\n", a.name)
	a.airportMediator.NotifyWaitingAircraft()
}

type helicopter struct {
	name            string
	airportMediator AirportMediator
}

func newHelicopter(name string, airportMediator AirportMediator) *helicopter {
	return &helicopter{name: name, airportMediator: airportMediator}
}

func (h *helicopter) ApproachAirport() {
	if !h.airportMediator.CanLandAirport(h) {
		fmt.Printf("机场繁忙，%s请继续等待\n", h.name)
		return
	}
	fmt.Printf("%s已降落\n", h.name)
}

func (h *helicopter) DepartAirport() {
	fmt.Printf("%s已经起飞离开机场;\n", h.name)
	h.airportMediator.NotifyWaitingAircraft()
}
