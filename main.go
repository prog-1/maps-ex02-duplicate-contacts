package main

import (
	"bufio"
	"fmt"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	PhoneBook := make(map[string]string)
	index := len(names)
	for i := 0; i != index; i++ {
		PhoneBook[names[i]] = numbers[i]
	}
	return PhoneBook
}

func main() {
	var names, numbers []string
	var i, size uint
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How many contacts do you want to enter?")
	fmt.Scan(&size)
	for i = 0; i != size; i++ {
		fmt.Print("Enter name #", i, " : ")
		text1, _ := reader.ReadString('\n')
		fmt.Print("Enter phone number #", i+1, " : ")
		text2, _ := reader.ReadString('\n')
		names = append(names, text1)
		numbers = append(numbers, text2)
	}
}
