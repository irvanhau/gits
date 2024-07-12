package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getCharWeight(char rune) int {
	alphabetMap := map[rune]int{
		'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8,
		'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16,
		'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24,
		'y': 25, 'z': 26,
	}

	value := alphabetMap[char]
	return value
}

func getSubstrWeight(str string) map[int]bool {
	weights := map[int]bool{}
	weight := 0
	for _, char := range str {
		weight += getCharWeight(char)
		weights[weight] = true
	}

	return weights
}

func checkString(weights map[int]bool, queries []int) []string {
	result := make([]string, len(queries))
	for i, char := range queries {
		if _, ok := weights[char]; ok {
			result[i] = "Yes"
		} else {
			result[i] = "No"
		}
	}
	return result
}

func main() {
	var inputString string
	var queries []int

	fmt.Println("Input string: ")
	fmt.Scanln(&inputString)
	fmt.Println("Input the queries, separated with whitespace")
	queries = GetInputSlice()

	weight := getSubstrWeight(inputString)
	result := checkString(weight, queries)

	fmt.Println("Output: ", result)
	fmt.Println("Queries: ", queries)
}

func numbers(s string) []int {
	var res []int
	for _, str := range strings.Fields(s) {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic("User Input Number Wrong")
		}
		res = append(res, i)
	}

	return res
}

func GetInputSlice() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return numbers(scanner.Text())
}
