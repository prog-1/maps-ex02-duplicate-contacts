package main

import (
	"fmt"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	pb := make(map[string]string)
	for i, name := range names {
		pb[name] = numbers[i]
	}
	return pb
}

func main() {
	fmt.Print("How many contacts do you want to enter?")
	var c int
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
