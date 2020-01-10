package main

import "fmt"

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImplementer interface {
	Send(text, to string)
}

type MessageSMS struct{}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS\n", text, to)
}

type MessageEmail struct{}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email\n", text, to)
}

type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}

func main() {
	m1 := NewCommonMessage(ViaSMS())
	m1.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via SMS

	m2 := NewCommonMessage(ViaEmail())
	m2.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via Email

	m3 := NewUrgencyMessage(ViaSMS())
	m3.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via SMS

	m4 := NewUrgencyMessage(ViaEmail())
	m4.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via Email
}
