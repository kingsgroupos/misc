package wlang

import "testing"

func TestConvertFormat(t *testing.T) {
	a := []struct {
		s1, s2 string
	}{
		{s1: "", s2: ""},
		{s1: "{{", s2: "{"},
		{s1: "}}", s2: "}"},
		{s1: "%", s2: "%%"},
		{s1: "{0}", s2: "%[1]v"},
		{s1: "{1}, {0}", s2: "%[2]v, %[1]v"},
		{s1: "%{1}, {{{0}}}", s2: "%%%[2]v, {%[1]v}"},
		{s1: "{0:0#}", s2: "%[1]v"},
		{s1: "{0:0#.00##}", s2: "%.4[1]f"},
		{s1: "{0:.00}", s2: "%.2[1]f"},
	}

	for i, v := range a {
		if r := convertFormat(v.s1); r != v.s2 {
			t.Fatalf("unexpected result. result: %q, expected: %q, i: %d", r, v.s2, i)
		}
	}
}
