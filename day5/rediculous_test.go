package main

import (
	"testing"
)

func Test_ShortWords_AreAlwaysNaughty(t *testing.T) {
	for _, str := range []string{"a", "aa", "aei"} {
		input := rediculousRules(str)
		expected := false
		actual := input.isNice()

		if expected != actual {
			t.Errorf("Expected %s to be nice, but wasn't", input)
		}
	}
}

func Test_Nice_ugknbfddgicrmopn(t *testing.T) {
	input := rediculousRules("ugknbfddgicrmopn")
	expected := true
	actual := input.isNice()

	if expected != actual {
		t.Errorf("Expected %s to be nice, but wasn't", input)
	}
}

func Test_Nice_aaa(t *testing.T) {
	input := rediculousRules("aaa")
	expected := true
	actual := input.isNice()

	if expected != actual {
		t.Errorf("Expected %s to be nice, but wasn't", input)
	}
}

func Test_Naughty_jchzalrnumimnmhp(t *testing.T) {
	input := rediculousRules("jchzalrnumimnmhp")
	expected := false
	actual := input.isNice()

	if expected != actual {
		t.Errorf("Expected %s to be naughty, but wasn't", input)
	}
}
