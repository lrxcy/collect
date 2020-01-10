package main

import "log"

type formaterFunc func(formater string) (string, error)

type formaterGenator interface {
	getfunc() formaterFunc
}

type Validator struct {
	choice string
}

func (v *Validator) getfunc() formaterFunc {
	switch v.choice {
	case "1":
		return Formater1
	case "2":
		return Formater2
	default:
		return nil

	}
	return nil
}

func Formater1(name string) (string, error) {
	return "foramter1 : " + name, nil
}

func Formater2(name string) (string, error) {
	return "foramter2 : " + name, nil
}

func main() {
	a := Validator{}
	a.choice = "1"
	f := a.getfunc()
	log.Println(f("tester1"))
	a.choice = "2"
	f2 := a.getfunc()
	log.Println(f2("tester2"))
}
