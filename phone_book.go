package main

import (
	"bufio"
	"fmt"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	m := make(map[string]string)
	for i := range names {
		m[names[i]] = numbers[i]
	}
	return m
}

func main() {
	var c int
	fmt.Scanln(&c)
	fmt.Println("How many contacts do you want to enter?")
	var name, number string
	var names, numbers []string
	for s := 1; s <= c; s++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter name #%v: ", s)
		scanner.Scan()
		name = scanner.Text()
		names = append(names, name)
		fmt.Printf("Enter phone number #%v: ", s)
		scanner.Scan()
		number = scanner.Text()
		numbers = append(numbers, number)
	}
	fmt.Println(createPhoneBook(names, numbers))
}
