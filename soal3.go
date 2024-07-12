package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "3943"
	k := 1
	res := highPalindromeWrap(s, k)

	fmt.Println(res)
}

func highPalindrome(s string, k int) string {
	n := len(s)

	if k < 0 || n == 1 {
		return "-1"
	}

	// Cek apakah string adalah palindrome
	if isPalindrome(s) {
		return s
	}

	// Memeriksa Palindrome tertinggi di setiap karakter secara rekursif
	var highest string
	for i := range s {
		newS := replaceChar(s, i, '9')
		candidate := highPalindrome(newS, k-1)
		if candidate > highest {
			highest = candidate
		}
	}
	return highest
}

// Function untuk mengecek apakah input adalah palindrome
func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func replaceChar(s string, i int, newChar byte) string {
	// Membuat string baru untuk menggantikan karakter pada index ke-i
	sb := []byte(s)
	sb[i] = newChar
	return string(sb)
}

func isValidNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func highPalindromeWrap(s string, k int) string {
	// Cek apakah input adalah angka (bukan string)
	if !isValidNumber(s) {
		return "-1"
	}

	// Handle error ketika input string kosong
	if s == "" {
		return "0"
	}

	return highPalindrome(s, k)
}
