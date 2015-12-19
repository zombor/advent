package main

import (
	"testing"
)

func Test_Better_Nice(t *testing.T) {
	for _, str := range []string{"xxyxx", "qjhvhtzxzqqjkmpb"} {
		input := betterRules(str)
		expected := true
		actual := input.isNice()

		if expected != actual {
			t.Errorf("Expected %s to be nice, but wasn't", input)
		}
	}
}

func Test_Better_Naughty(t *testing.T) {
	for _, str := range []string{"uurcxstgmygtbstg", "ieodomkazucvgmuy"} {
		input := betterRules(str)
		expected := false
		actual := input.isNice()

		if expected != actual {
			t.Errorf("Expected %s to be naughty, but wasn't", input)
		}
	}
}

func Test_Better_passing_hasSubPairTwice(t *testing.T) {
	for _, str := range []string{"aabcdefgaa", "xyxy"} {
		input := betterRules(str)
		expected := true
		actual := input.hasSubpairTwice()

		if expected != actual {
			t.Errorf("Expected %s to have subpair twice, but didn't", input)
		}
	}
}

func Test_Better_failing_hasSubPairTwice(t *testing.T) {
	input := betterRules("aaa")
	expected := false
	actual := input.hasSubpairTwice()

	if expected != actual {
		t.Errorf("Expected %s to not have subpair twice, but did", input)
	}
}

func Test_Better_passing_hasRepeatingNonConsecutiveLetters(t *testing.T) {
	for _, str := range []string{"xyx", "abcdefeghi", "aaa"} {
		input := betterRules(str)
		expected := true
		actual := input.hasRepeatingNonConsecutiveLetters()

		if expected != actual {
			t.Errorf("Expected %s to have repeating nonconsecutive letters, but didn't", input)
		}
	}
}
