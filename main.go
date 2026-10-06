package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a number ")
	scanner.Scan()
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("You have to enter a number")
	} else if n > 7 {
		fmt.Println("Hello")
	}
	fmt.Print("Enter a name ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	if name == "John" {
		fmt.Println("Hello, John")
	} else {
		fmt.Println("There is no such name")
	}
	fmt.Print("Enter numbers separated by spaces")
	scanner.Scan()
	var nums []int
	for _, f := range strings.Fields(scanner.Text()) {
		v, err := strconv.Atoi(f)
		if err != nil {
			fmt.Print("Text you wrote is not a number, skipped")
			continue
		}
		nums = append(nums, v)
	}

	fmt.Print("Multiples of 3:")
	for _, v := range nums {
		if v%3 == 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Print("Write down brackets ")
	scanner.Scan()
	if isValid(strings.TrimSpace(scanner.Text())) {
		fmt.Println("The string is correct")
	} else {
		fmt.Println("The string is not correct")
	}
}

func isValid(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{'}
	var runes []rune
	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			runes = append(runes, ch)
		case ')', ']', '}':
			if len(runes) == 0 || runes[len(runes)-1] != pairs[ch] {
				return false
			}
			runes = runes[:len(runes)-1]
		}
	}
	return len(runes) == 0
}
