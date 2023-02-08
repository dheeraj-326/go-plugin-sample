package modules

type Adder interface {
	Add(x, y int) int
}

type adder struct {
}

func (a *adder) Add(x, y int) int {
	return x + y
}

func NewAdder() interface{} {
	return &adder{}
}
