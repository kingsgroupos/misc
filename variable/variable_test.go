package variable

import "testing"

func TestToCamel(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{"", ""},
		{"test", "test"},
		{"Test", "test"},
		{"TEST", "test"},
		{"testCase", "testCase"},
		{"TestCase", "testCase"},
		{"TESTCase", "testCase"},
		{"test_case", "testCase"},
		{"test_Case", "testCase"},
		{"Test_case", "testCase"},
		{"Test_Case", "testCase"},
		{"TEST_case", "testCase"},
		{"TEST_Case", "testCase"},
		{"test1", "test1"},
		{"Test1", "test1"},
		{"TEST1", "test1"},
		{"test1x", "test1X"},
		{"test_1", "test1"},
		{"Test_1", "test1"},
		{"TEST_1", "test1"},
		{"test12", "test12"},
		{"Test12", "test12"},
		{"TEST12", "test12"},
		{"test_12", "test12"},
		{"Test_12", "test12"},
		{"TEST_12", "test12"},
		{"test123", "test_123"},
		{"Test123", "test_123"},
		{"TEST123", "test_123"},
		{"test_123", "test_123"},
		{"Test_123", "test_123"},
		{"TEST_123", "test_123"},
		{"_", ""},
		{"中文", ""},
		{"_a", "a"},
		{"a_", "a"},
		{"_a_", "a"},
		{"__a", "a"},
		{"a__", "a"},
		{"__a__", "a"},
		{"__a__b", "aB"},
		{"__a__bc__", "aBc"},
		{"__a__bC__", "aBC"},
		{"__a__Bc__", "aBc"},
		{"__a__BC__", "aBC"},
		{"1", ""},
		{"_1", ""},
		{"1a", "a"},
		{"_1a", "a"},
	}

	for i, c := range cases {
		if o := ToCamel(c.input); o != c.output {
			t.Fatalf("ToCamel does not work as expected. i: %d, input: %s, expected: %s, actual: %s",
				i, c.input, c.output, o)
		}
	}
}
