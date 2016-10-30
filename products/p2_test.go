package products

import (
	"testing"
)

func TestNewP2(t *testing.T) {
	p2 := NewP2()
	if p2.Name() != "P2" {
		t.Error("Error occurred")
	}
}