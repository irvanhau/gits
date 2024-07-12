package main

import "fmt"

func main() {
	var input string

	fmt.Println("Input Bracket, seperated by whitespace or not")
	fmt.Scan(&input)

	res := isBalanced(input)

	fmt.Println(res)
}

func isMatching(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '{':
		return close == '}'
	case '[':
		return close == ']'
	default:
		return false
	}
}

func isBalanced(str string) string {
	stack := []rune{}
	for _, char := range str {
		// fmt.Println("Ini Stack: ", len(stack), ", Ini Char: ", string(char))
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || !isMatching(stack[len(stack)-1], char) {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	}

	return "NO"
}
