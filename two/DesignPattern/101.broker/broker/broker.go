package broker

import (
	"github.com/jimweng/DesignPattern/101.broker/broker/packagea"
	"github.com/jimweng/DesignPattern/101.broker/broker/packageb"
)

type Broker interface {
	Execute()
}

func NewA() Broker {
	// return ABroker(new(packagea.Delegate))
	// return new(packagea.Delegate)
	return &packagea.Delegate{}
}

func NewB() Broker {
	return &packageb.Delegate{}
}
