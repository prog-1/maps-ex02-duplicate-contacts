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
	var count int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("How many contacts do you want to enter?")
	scanner.Scan()
	bcount := scanner.Text()
	fmt.Sscan(bcount, &count)
	names := make([]string, count)
	numbers := make([]string, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Enter name #%v: ", i+1)
		scanner.Scan()
		names[i] = scanner.Text()
		fmt.Printf("Enter phone number #%v: ", i+1)
		scanner.Scan()
		numbers[i] = scanner.Text()
	}
	fmt.Println(createPhoneBook(names, numbers))
}
