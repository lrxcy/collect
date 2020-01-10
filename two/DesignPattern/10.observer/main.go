package main

import "fmt"

// Observer定義方法Update更新Register
type Observer interface {
	Update(*Register)
}

// Register內部提供observers，紀錄這些被觀察者，且每個觀察者都具有Update的方法
type Register struct {
	observers []Observer
	context   string
}

// NewRegister() 定義了創建該觀察列表的方法，回傳一個Register，方便讓後面的Observer做註冊
func NewRegister() *Register {
	return &Register{
		observers: make([]Observer, 0),
	}
}

// Attach 提供後面新增的Observer做註冊的方法
func (s *Register) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

// notify 提供一個對於所有在Register內Observer操作的方法
func (s *Register) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

// UpdateContext只是一個將notify包裝起來的方法，旨在更新所有Observer下的物件
func (s *Register) UpdateContext(context string) {
	s.context = context
	s.notify()
}

func main() {
	Register := NewRegister()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	Register.Attach(reader1)
	Register.Attach(reader2)
	Register.Attach(reader3)

	Register.UpdateContext("observer mode")
}

// Reader: 定義一個可以被加入到觀察list的物件
type Reader struct {
	name string
}

// NewReader: 定義如何生成該觀察物件的方法
func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

// Update: 定義更新的方法
func (r *Reader) Update(s *Register) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}
