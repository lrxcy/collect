package decorator

type Component interface {
	ReturnInt() int
	ReturnString() string
}

type Validator struct {
	Component
}

func WrapValidator(c Component) Component {
	return &Validator{
		Component: c,
	}
}

func (w *Validator) ReturnInt() int {
	if w.Component.ReturnInt() >= 20 {
		if w.Component.ReturnInt() >= 30 {
			return w.Component.ReturnInt() - 10
		} else {
			return w.Component.ReturnInt() + 1
		}
	}
	return w.Component.ReturnInt()
}

func (w *Validator) ReturnString() string {
	return w.Component.ReturnString() + " test"
}
