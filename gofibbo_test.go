package gofibbo

import "testing"

func TestFibbo(t *testing.T) {
	value := Fibbo(44)
	want := int64(701408733)
	if value != want {
		t.Errorf("Fibbo(10) != %v, want %v", value, want)
	}
}
