package domain

type Adder interface {
	Add(x, y int) int
}

type Divider interface {
	Divide(x, y int) float64
}
