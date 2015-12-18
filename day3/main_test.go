package main

import (
	"testing"
)

func Test_MoveRight_CountsTwo(t *testing.T) {
	input := []string{">"}
	expected := 2
	actual := countDeliveredHouses(input)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_MoveRightTwice_CountsThree(t *testing.T) {
	input := []string{">", ">"}
	expected := 3
	actual := countDeliveredHouses(input)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_MoveInBox_CountsFour(t *testing.T) {
	input := []string{"^", ">", "v", "<"}
	expected := 4
	actual := countDeliveredHouses(input)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_MoveUpAndDownABunch_CountsTwo(t *testing.T) {
	input := []string{"^", "v", "^", "v", "^", "v", "^", "v", "^", "v"}
	expected := 2
	actual := countDeliveredHouses(input)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}
