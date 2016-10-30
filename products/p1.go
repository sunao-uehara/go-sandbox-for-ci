package products

import (
	"fmt"
)

type P1 struct{}
func NewP1() *P1 {
	return &P1{}
}

func (p1 *P1) Name() string {
	fmt.Println("P1 - Name()")
	return "P1"
}
