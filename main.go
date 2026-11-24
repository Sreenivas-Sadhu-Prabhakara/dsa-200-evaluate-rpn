package main

import (
	"fmt"
	"strconv"
)

// EvalRPN evaluates a Reverse Polish Notation expression: operands push,
// operators pop two and push the result. ok is false for malformed input
// (too few operands, leftovers, unknown token, division by zero).
// Division truncates toward zero, as Go's integer division already does.
func EvalRPN(tokens []string) (int, bool) {
	stack := []int{}
	for _, tok := range tokens {
		switch tok {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, false
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var v int
			switch tok {
			case "+":
				v = a + b
			case "-":
				v = a - b
			case "*":
				v = a * b
			default:
				if b == 0 {
					return 0, false
				}
				v = a / b
			}
			stack = append(stack, v)
		default:
			n, err := strconv.Atoi(tok)
			if err != nil {
				return 0, false
			}
			stack = append(stack, n)
		}
	}
	if len(stack) != 1 {
		return 0, false
	}
	return stack[0], true
}

func main() {
	fmt.Println(EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
