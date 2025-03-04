package something_test

import (
	something "7-7"
	"testing"
)

func FuzzDoSomething(f *testing.F) {
	f.Add("test&&&")
	f.Add("hello world")
	f.Add("valid_input")

	f.Fuzz(func(t *testing.T, s string) {
		err := something.DoSomething(s)
		if err != nil && s != "test&&&" {
			t.Errorf("unexpected error for input %q: %v", s, err)
		}
	})
}
