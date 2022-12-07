package observer

import "fmt"

type MsgType int

const (
	ConsumeType MsgType = iota
	BillType
	ExpireType
)

type CreditCard struct {
	holder          string
	consumeSum      float32
	subscriberGroup map[MsgType][]Subscriber
}

func NewCreditCard(holder string) *CreditCard {
	return &CreditCard{holder: holder, subscriberGroup: make(map[MsgType][]Subscriber)}
}

func (c *CreditCard) Subscribe(subscriber Subscriber, msgTypes ...MsgType) {
	for _, msgType := range msgTypes {
		c.subscriberGroup[msgType] = append(c.subscriberGroup[msgType], subscriber)
	}
}

func (c *CreditCard) Unsubscribe(subscriber Subscriber, msgTypes ...MsgType) {
	for _, msgType := range msgTypes {
		if subs, ok := c.subscriberGroup[msgType]; ok {
			c.subscriberGroup[msgType] = removeSubscriber(subs, subscriber)
		}
	}
}

func removeSubscriber(subscribers []Subscriber, toRemove Subscriber) []Subscriber {
	length := len(subscribers)
	for i, subscriber := range subscribers {
		if subscriber.Name() == toRemove.Name() {
			subscribers[length-1], subscribers[i] = subscribers[i], subscribers[length-1]
			return subscribers[:length-1]
		}
	}
	return subscribers
}

func (c *CreditCard) Consumer(money float32) {
	c.consumeSum += money
	c.notify(ConsumeType, fmt.Sprintf("持卡人%s,本月账单金额%.2f元\n", c.holder, c.consumeSum))
}

func (c *CreditCard) SendBill() {
	c.notify(BillType, fmt.Sprintf("尊敬的持卡人%s,您本月账单已出，消费总额%.2f元;\n", c.holder, c.consumeSum))
}

func (c *CreditCard) Expire() {
	c.notify(ExpireType, fmt.Sprintf("尊敬的持卡人%s,您本月账单已逾期，请及时还款，总额%.2f元;\n", c.holder, c.consumeSum))
}

func (c *CreditCard) notify(msgType MsgType, message string) {
	if subs, ok := c.subscriberGroup[msgType]; ok {
		for _, sub := range subs {
			sub.Update(message)
		}
	}
}
