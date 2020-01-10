package main

import "log"

/*
目標: 手機接上轉接器，轉接器插上插頭
1. 手機到轉接器的充電方法
2. 轉接器到插頭的介面調用
3. 插頭如何幫手機充電
*/

// ----- define the plug ----
type PlugImple interface {
	Charging() string
}

type Plug struct{}

func (*Plug) Charging() string {
	return "device is charging"
}

func NewPlug() PlugImple {
	return &Plug{}
}

// ---define another Plug Implement with inhereitance of PlugImple
/*
	當如果需要定義其他的插頭時，可以考慮直接繼承PlugImple，
	並且繼承相同的方法
*/
type PlugImple2 interface {
	PlugImple
}

type Plug2 struct{}

func (*Plug2) Charging() string {
	return "device is not chargin"
}

func NewPlug2() PlugImple2 {
	return &Plug2{}
}

// ------ define the adapter---------
type PowerAdapteeImpl interface {
	Charge() string
}

type PowerAdaptee struct {
	PlugImple
}

// 這邊以`a`表示該PowerAdaptee，調度a下面的方法Charging()來回傳字串
func (a *PowerAdaptee) Charge() string {
	return a.Charging()
}

func NewPowerAdaptee(pluginImple PlugImple) PowerAdapteeImpl {
	return &PowerAdaptee{
		PlugImple: pluginImple,
	}
}

func main() {
	f := NewPlug()
	ff := NewPowerAdaptee(f)
	log.Println(ff.Charge())

	// 實作第二個插座
	f2 := NewPlug2()
	ff2 := NewPowerAdaptee(f2)
	log.Println(ff2.Charge())
}
