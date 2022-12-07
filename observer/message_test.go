package observer

import "testing"

func TestObserver(t *testing.T) {
	creditCard := NewCreditCard("Bob")
	creditCard.Subscribe(new(shortMessage), ConsumeType, ExpireType)
	creditCard.Subscribe(new(email), BillType, ExpireType)
	creditCard.Subscribe(new(telephone), ExpireType)

	creditCard.Consumer(500)
	creditCard.Consumer(300)
	creditCard.SendBill()
	creditCard.Expire()
	creditCard.Unsubscribe(new(email), ExpireType)
	creditCard.Unsubscribe(new(shortMessage), ExpireType)
	creditCard.Consumer(100)
	creditCard.Expire()
}
