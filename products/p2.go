package products

import (
	"fmt"
)

type P2 struct{}
func NewP2() *P2 {
	return &P2{}
}

func (p2 *P2) Name() string {
	fmt.Println("P2 - Name()")
	return "P2"
}
