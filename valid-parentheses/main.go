package main

import "fmt"

func main() {

	fmt.Println(isValid("()[]{}"))

}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	//pairs['{']='}'

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []rune

	// if (,{,[ look in pairs
	//if yes append to stack
	// if not looking stack[i-1] -- if yes >> true

	for _, v := range s {
		if _, ok := pairs[v]; ok {
			stack = append(stack, v)
			// looking for stack
			// v = )
			// if pairs['('] = )
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != v {
			return false
			// stack = empty[]rune
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
