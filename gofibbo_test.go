package gofibbo

import "testing"

func TestFibbo(t *testing.T) {
	value := Fibbo(10)
	if value != 55 {
		t.Errorf("Fibbo(10) != %v, want 55", Fibbo(10))
	}
}
