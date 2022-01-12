package main

import (
	"fmt"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	x := make(map[string]string)
	for a, b := range names {
		x[b] = numbers[a]
	}
	return x
}

func main() {
	var c int
	fmt.Print("How many contacts do you want to enter?")
	fmt.Scan(&c)
	var names, numbers []string
	var scan string
	for y := 1; y <= c; y++ {
		fmt.Printf("Enter name #%v:", y)
		fmt.Scan(&scan)
		names = append(names, scan)
		fmt.Printf("Enter phone number #%v:", y)
		fmt.Scan(&scan)
		numbers = append(numbers, scan)
	}
	fmt.Println(createPhoneBook(names, numbers))
}
