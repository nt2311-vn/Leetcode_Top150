package validparentheses

import "testing"

func TestValidOne(t *testing.T) {
	s := "()"
	expected := true

	actual := isValid(s)
	if actual != expected {
		t.Errorf("Failed on 1st test expected %v but got %v\n", expected, actual)
	}
}

func TestValidTwo(t *testing.T) {
	s := "()[]{}"
	expected := true

	actual := isValid(s)
	if actual != expected {
		t.Errorf("Failed on 2nd test expected %v but got %v\n", expected, actual)
	}
}

func TestValidThree(t *testing.T) {
	s := "(]"
	expected := false

	actual := isValid(s)
	if actual != expected {
		t.Errorf("Failed on 2nd test expected %v but got %v\n", expected, actual)
	}
}
