package main

import (
	"bufio"
	"fmt"
	"os"
)

func CreatePhoneBook(names []string, numbers []string) map[string]string {
	m := make(map[string]string)
	for i, key := range names {
		m[key] = numbers[i]
	}
	return m
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("This program creates a phone book storing names and phone numbers of your contacts.")
	fmt.Print("How many contacts do you want to enter?")
	var contacts int
	fmt.Scan(&contacts)
	var name, number string
	var names, numbers []string
	for i := 1; i <= contacts; i++ {
		fmt.Printf("Enter name #%v: ", i)
		scanner.Scan()
		name = scanner.Text()
		names = append(names, name)
		fmt.Printf("Enter phone number #%v: ", i)
		scanner.Scan()
		number = scanner.Text()
		numbers = append(numbers, number)
	}
	fmt.Println(CreatePhoneBook(names, numbers))

}
