package main

import (
	"bufio"
	"fmt"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	PhoneBook := make(map[string]string)
	for i, key := range names {
		PhoneBook[key] = numbers[i]
	}
	return PhoneBook
}

func main() {
	fmt.Println(`Program that creates a phone book storing names
and phone numbers of your contacts.`)
	fmt.Println("How many contacts do you want to enter?")
	var count int
	fmt.Scanln(&count)
	for i := 0; i < count; i++ {
		var names, numbers []string
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter name #%v: ", i+1)
		var s1 string
		scanner.Scan()
		s1 = scanner.Text()
		fmt.Printf("Enter phone number #%v: ", i+1)
		var s2 string
		scanner.Scan()
		s2 = scanner.Text()
		names = append(names, s1)
		numbers = append(numbers, s2)
		fmt.Println(createPhoneBook(names, numbers))
	}
}

// program isn't full. Output should be at the end with all names and numbers (map[aa:111]). But now it is after each "contact"
