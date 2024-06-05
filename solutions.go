package main

import (
	"os"

	print "github.com/01-edu/z01"
)

func RevParam() {
	result := "abcdefghijklmnopqrstuvwxyz"

	f := []rune{}
	for i, char := range result {
		// uppercase the elements whose index is odd
		if i%2 == 0 {
			f = append(f, toupper(char))
		} else if i%2 == 1 {
			f = append(f, char)
		}
	}

	// reverse the string
	for i, j := 0, len(f)-1; i < len(f)/2; i, j = i+1, j-1 {
		f[i], f[j] = f[j], f[i]
	}
	for _, char := range f {
		print.PrintRune(char)
	}
	print.PrintRune('\n')
}

func toupper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		r = r - 32
	}
	return r
}

func FirstParam() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	result := args[0]
	for _, char := range result {
		print.PrintRune(char)
	}
	print.PrintRune('\n')
}

func LastParam() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	result := args[len(args)-1]
	for _, char := range result {
		print.PrintRune(char)
	}
	print.PrintRune('\n')
}

func itoa(num int) string {
	// if the number is 0
	if num == 0 {
		return "0"
	}

	// check if the number is negative
	var isNegative bool = false
	if num < 0 {
		isNegative = true
		num = -num
	}

	result := []rune{}
	for num > 0 {
		lastNum := num % 10
		result = append(result, rune(48+lastNum))
		num = num / 10
	}

	if isNegative {
		result = append(result, '-')
	}

	reverse := func(slice []rune) []rune {
		for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
			slice[i], slice[j] = slice[j], slice[i]
		}
		return slice
	}

	result = reverse(result)
	return string(result)
}

func ParamCount() {
	args := os.Args[1:]
	n := len(args)
	if n < 1 {
		return
	}

	str := itoa(n)
	for _, num := range str {
		print.PrintRune(num)
	}
	print.PrintRune('\n')
}

func CountDown() {
	num := "0123456789"
	result := []rune(num)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	for _, num := range result {
		print.PrintRune(num)
	}
	print.PrintRune('\n')
}

func Atoi(s string) int {
	if s == "0" {
		return 0
	}
	result := 0

	isNegative := false
	if s[0] == '-' {
		s = Pop(s, 0)
		isNegative = true
	} else if s[0] == '+' {
		s = Pop(s, 0)
		isNegative = false
	}

	for _, num_str := range s {
		result = result*10 + int(rune(num_str-48))
	}

	if isNegative {
		result = -result
	}

	return result
}

func InsertAt(s string, index int, char rune) string {
	// check the index out of bounds
	if index < 0 || index >= len(s) {
		return ""
	}
	runes := []rune(s)
	// insert at a certain index
	result := append(runes[index:], append([]rune{char}, runes[index:]...)...)
	return string(result)
}

func Pop(s string, index int) string {
	runes := []rune(s)
	result := append(runes[:index], runes[index+1:]...)
	return string(result)
}

func WordMatch() string {
	args := os.Args[1:]

	n := len(args)

	if n < 2 {
		return ""
	}
	firstword := args[0]
	secondword := args[1]
	i, j := 0, 0
	for i < len(firstword) && j < len(secondword) {
		if firstword[i] == secondword[j] {
			i++
		}
		j++

		if i == len(firstword) {
			return firstword
		}
	}
	return ""
}

func LastRune(s string) rune {
	runeSlice := []rune(s)
	return runeSlice[len([]rune(s))-1]
}

func ReduceInt(a []int, f func(int, int) int) int {
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	return result
}
