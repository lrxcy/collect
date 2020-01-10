package main

import "log"

// Cloneable 為一個interface回傳interface的設計
type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeManager) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

func (p *PrototypeManager) Get(name string) Cloneable {
	return p.prototypes[name]
}

func main() {

	manager := NewPrototypeManager()

	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)

	t2 := t1.Clone()
	if t1 != t2 {
		log.Println("success clone a prototype")
	}

	c := manager.Get("t1").Clone()

	tt1 := c.(*Type1)
	log.Println(tt1.name)

}

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}
