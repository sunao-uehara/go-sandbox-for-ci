package products

import (
	"testing"
)

func TestRegister(t *testing.T) {
	// redeclare p1
	Register("p1", NewP1())

	// sad path passing nil product
	err := Register("p1", nil)
	if err.Error() != "Product factory p1 does not exist" {
		t.Error("should occurr error")
	}
}

func TestGetProduct(t *testing.T) {
	// happay path
	p, err := GetProduct("p1")
	if p.Name() != "P1" {
		t.Error("invalid product")
	}

	// sad path
	_, err = GetProduct("p3")
	if err == nil {
		t.Error("")
	}
}
