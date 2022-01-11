package main

import (
	"bufio"
	"fmt"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	m := make(map[string]string)
	for i, key := range names {
		m[key] = numbers[i]
	}
	return m
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var c int
	fmt.Println("How many contacts do you want to enter?")
	fmt.Scan(&c)
	var name, number string
	var names, numbers []string
	for i := 1; i <= c; i++ {
		fmt.Printf("Enter name #%v: ", i)
		scanner.Scan()
		name = scanner.Text()
		names = append(names, name)
		fmt.Printf("Enter phone number #%v: ", i)
		scanner.Scan()
		number = scanner.Text()
		numbers = append(numbers, number)
	}
	fmt.Println(createPhoneBook(names, numbers))
}
