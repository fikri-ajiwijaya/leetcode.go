package problem_0003_test

import (
	"testing"

	. "github.com/fikri-ajiwijaya/leetcode.go/problem_0003"
)

func HasDuplicate_test(t *testing.T, s string, expected bool) {
	result := HasDuplicate(s)
	if expected != result {
		t.Errorf("s = %#v", s)
		t.Errorf("expected = %#v", expected)
		t.Errorf("result = HasDuplicate(s)")
		t.Errorf("result is %#v", result)
	}
}

func LengthOfLongestSubstring_test(t *testing.T, s string, expected int) {
	result := LengthOfLongestSubstring(s)
	if expected != result {
		t.Errorf("s = %#v", s)
		t.Errorf("expected = %#v", expected)
		t.Errorf("result = LengthOfLongestSubstring(s)")
		t.Errorf("result is %#v", result)
	}
}

func TestHasDuplicatev1(t *testing.T) {
	HasDuplicate_test(t, "aba", true)
}

func TestHasDuplicatev2(t *testing.T) {
	HasDuplicate_test(t, "abc", false)
}

func TestExample1(t *testing.T) {
	LengthOfLongestSubstring_test(t, "abcabcbb", 3)
}

func TestExample2(t *testing.T) {
	LengthOfLongestSubstring_test(t, "bbbbb", 1)
}

func TestExample3(t *testing.T) {
	LengthOfLongestSubstring_test(t, "pwwkew", 3)
}
