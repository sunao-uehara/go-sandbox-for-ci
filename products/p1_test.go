package products

import (
	"testing"
)

func TestNewP1(t *testing.T) {
	p1 := NewP1()
	if p1.Name() != "P1" {
		t.Error("Error occurred")
	}
}