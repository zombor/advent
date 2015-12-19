package main

import (
	"testing"
)

func Test_ShortWords_AreAlwaysNaughty(t *testing.T) {
	for _, str := range []string{"a", "aa", "aei"} {
		input := str
		expected := false
		actual := isNice(input)

		if expected != actual {
			t.Errorf("Expected %s to be nice, but wasn't", input)
		}
	}
}

func Test_Nice_ugknbfddgicrmopn(t *testing.T) {
	input := "ugknbfddgicrmopn"
	expected := true
	actual := isNice(input)

	if expected != actual {
		t.Errorf("Expected %s to be nice, but wasn't", input)
	}
}

func Test_Nice_aaa(t *testing.T) {
	input := "aaa"
	expected := true
	actual := isNice(input)

	if expected != actual {
		t.Errorf("Expected %s to be nice, but wasn't", input)
	}
}

func Test_Naughty_jchzalrnumimnmhp(t *testing.T) {
	input := "jchzalrnumimnmhp"
	expected := false
	actual := isNice(input)

	if expected != actual {
		t.Errorf("Expected %s to be naughty, but wasn't", input)
	}
}
