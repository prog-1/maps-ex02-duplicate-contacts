package main

import (
	"fmt"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	l := make(map[string]string)
	for i, kt := range names {
		l[kt] = numbers[i]
	}
	return l
}

func main() {
	var c int
	fmt.Print("How many contacts do you want to enter?")
	fmt.Scan(&c)
	var names, numbers []string
	var tmp string
	for i := 1; i <= c; i++ {
		fmt.Printf("Enter name #%v:", i)
		fmt.Scan(&tmp)
		names = append(names, tmp)
		fmt.Printf("Enter phone number #%v:", i)
		fmt.Scan(&tmp)
		numbers = append(numbers, tmp)
	}
	fmt.Print(createPhoneBook(names, numbers))
}
