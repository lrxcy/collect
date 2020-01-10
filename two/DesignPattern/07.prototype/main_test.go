package main

import "testing"

var manager *PrototypeManager

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func TestClone(t *testing.T) {
	t2 := manager.Get("t2")

	tt2 := t2.Clone()

	if t2 == tt2 {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t2").Clone()

	t2 := c.(*Type2)
	if t2.name != "type2" {
		t.Fatal("error")
	}

}

func init() {
	manager = NewPrototypeManager()

	t2 := &Type2{
		name: "type2",
	}
	manager.Set("t2", t2)
}
