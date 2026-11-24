package main

import "testing"

func TestEvalRPN(t *testing.T) {
	cases := []struct {
		in   []string
		want int
		ok   bool
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9, true},
		{[]string{"4", "13", "5", "/", "+"}, 6, true},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22, true},
		{[]string{"7"}, 7, true},
		{[]string{"6", "-132", "/"}, 0, true}, // truncation toward zero
		{[]string{"+", "1"}, 0, false},        // operator before operands
		{[]string{"1", "2"}, 0, false},        // leftovers
		{[]string{"1", "0", "/"}, 0, false},   // division by zero
		{[]string{"x"}, 0, false},             // unknown token
	}
	for _, c := range cases {
		got, ok := EvalRPN(c.in)
		if ok != c.ok || (ok && got != c.want) {
			t.Errorf("EvalRPN(%v) = (%d,%v), want (%d,%v)", c.in, got, ok, c.want, c.ok)
		}
	}
}
