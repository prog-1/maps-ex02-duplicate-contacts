package main

import (
	"bufio"
	"fmt"
	"os"
)

func CreatePhoneBook(names []string, numbers []string) map[string]string {
	x := len(names)
	phonebook := make(map[string]string)
	for i := 0; x > i; i++ {
		phonebook[names[i]] = numbers[i]
	}
	return phonebook
}

func main() {
	var names, numbers []string
	var size uint
	fmt.Print("How many contacts do you want to enter: ")
	fmt.Scanln(&size)
	for n := 1; size > 0; {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter name №%v: ", n)
		scanner.Scan()
		name := scanner.Text()
		names = append(names, name)
		fmt.Printf("Enter number №%v: ", n)
		scanner.Scan()
		number := scanner.Text()
		numbers = append(numbers, number)
		n++
		size--
	}
	fmt.Println(CreatePhoneBook(names, numbers))
}
