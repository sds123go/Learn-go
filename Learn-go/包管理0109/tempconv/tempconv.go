package tempconv

import "fmt"

type Sheshidu float64
type HuaShidu float64

const (
	AbsoluteZero Sheshidu = -273.15
	FreezingC    Sheshidu = 0
	BoilingC     Sheshidu = 100
)

func (s Sheshidu) String() string {
	return fmt.Sprintf("%g°C", s)
}
func (h HuaShidu) String() string {
	return fmt.Sprintf("%g°C", h)
}
