package main

import (
	"bufio"
	"fmt"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	m := make(map[string]string)
	a := len(names)
	fmt.Println(a)
	for i := 0; i < a; i++ {
		m[names[i]] = numbers[i]
	}
	return m
}
func main() {
	var a int
	fmt.Println("How many contacts do you want to enter?")
	fmt.Scanln(&a)
	var s string
	var names, numbers []string
	for i := 0; i < a; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter name #%v: ", i+1)
		scanner.Scan()
		s = scanner.Text()
		names = append(names, s)
		fmt.Printf("Enter phone number #%v: ", i+1)
		scanner.Scan()
		s = scanner.Text()
		numbers = append(numbers, s)
	}
	fmt.Println(createPhoneBook(names, numbers))
}
