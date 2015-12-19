package main

import (
	"reflect"
	"testing"
)

func Test_MoveRight(t *testing.T) {
	input := []string{">"}
	expected := 2
	actual := countDeliveredHouses(deliveredHouses(input).positions)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_MoveRightTwice(t *testing.T) {
	input := []string{">", ">"}
	expected := 3
	actual := countDeliveredHouses(deliveredHouses(input).positions)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_MoveInBox(t *testing.T) {
	input := []string{"^", ">", "v", "<"}
	expected := 4
	actual := countDeliveredHouses(deliveredHouses(input).positions)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_MoveUpAndDownABunch(t *testing.T) {
	input := []string{"^", "v", "^", "v", "^", "v", "^", "v", "^", "v"}
	expected := 2
	actual := countDeliveredHouses(deliveredHouses(input).positions)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_SantaAndRobotSanta_MoveInBox(t *testing.T) {
	input := []string{"^", ">", "v", "<"}
	expected := 3
	actual := countCombinedDeliveredHouses(input)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_SantaAndRobotSanta_MoveInBoxTwice(t *testing.T) {
	input := []string{"^", ">", "v", "<", "^", ">", "v", "<"}
	expected := 3
	actual := countCombinedDeliveredHouses(input)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_SantaAndRoboSanta_MoveUpAndDown(t *testing.T) {
	input := []string{"^", "v"}
	expected := 3
	actual := countCombinedDeliveredHouses(input)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_SantaAndRoboSanta_MoveUpAndDownABunch(t *testing.T) {
	input := []string{"^", "v", "^", "v", "^", "v", "^", "v", "^", "v"}
	expected := 11
	actual := countCombinedDeliveredHouses(input)

	if actual != expected {
		t.Errorf("Expected %#v to equal %d, but got %d", input, expected, actual)
	}
}

func Test_splitEven(t *testing.T) {
	input := []string{"0", "1", "2", "3", "4"}
	expected := []string{"0", "2", "4"}
	actual := splitEven(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v to equal %#v, but got %#v", input, expected, actual)
	}
}

func Test_splitOdd(t *testing.T) {
	input := []string{"0", "1", "2", "3", "4"}
	expected := []string{"1", "3"}
	actual := splitOdd(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v to equal %#v, but got %#v", input, expected, actual)
	}
}
