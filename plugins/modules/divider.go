package modules

type Divider interface {
	Divide(x, y int) float64
}

type divider struct {
}

func (a *divider) Divide(x, y int) float64 {
	return float64(x) / float64(y)
}

func NewDivider() interface{} {
	return &divider{}
}
