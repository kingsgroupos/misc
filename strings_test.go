package misc

import (
	"testing"
)

func TestSplit(t *testing.T) {
	inputs := []string{
		"hello world 123",
		"   hello world 123",
		"hello world 123   ",
		"hello  world  123",
	}

	for _, str := range inputs {
		a := Split(str, " ")
		if a[0] != "hello" || a[1] != "world" || a[2] != "123" {
			t.Fatalf("something is wrong with misc.Split. input: %q, result: %s", str, ToJSON(a))
		}
	}
}

func TestUnescapeUnicode(t *testing.T) {
	str1 := `Journey to the \u300EWest\u300F`
	if UnescapeUnicode(str1) != "Journey to the 『West』" {
		t.Fatalf("failed to unquote the following string: %s", str1)
	}

	str2 := `Journey to the \u300EWest\u300F - Wu Chengen`
	if UnescapeUnicode(str2) != "Journey to the 『West』 - Wu Chengen" {
		t.Fatalf("failed to unquote the following string: %s", str2)
	}
}

func TestUCFirst(t *testing.T) {
	strs1 := []string{
		"a",
		"A",
		"abc",
		"Abc",
		"程序员",
		"a程",
		"A程",
	}
	strs2 := []string{
		"A",
		"A",
		"Abc",
		"Abc",
		"程序员",
		"A程",
		"A程",
	}

	if len(strs1) != len(strs2) {
		t.Fatalf("strs1 and strs2 are not the same length. len1: %d, len2: %d", len(strs1), len(strs2))
	}
	for i := 0; i < len(strs1); i++ {
		if UCFirst(strs1[i]) != strs2[i] {
			t.Fatalf("UCFirst does not work. i: %d, s1: %s, s2: %s", i, strs1[i], strs2[i])
		}
	}
}

func TestLCFirst(t *testing.T) {
	strs1 := []string{
		"a",
		"A",
		"abc",
		"Abc",
		"程序员",
		"a程",
		"A程",
	}
	strs2 := []string{
		"a",
		"a",
		"abc",
		"abc",
		"程序员",
		"a程",
		"a程",
	}

	if len(strs1) != len(strs2) {
		t.Fatalf("strs1 and strs2 are not the same length. len1: %d, len2: %d", len(strs1), len(strs2))
	}
	for i := 0; i < len(strs1); i++ {
		if LCFirst(strs1[i]) != strs2[i] {
			t.Fatalf("LCFirst does not work. i: %d, s1: %s, s2: %s", i, strs1[i], strs2[i])
		}
	}
}
