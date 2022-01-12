package main

import (
	"bufio"
	"fmt"
	"os"
)

func PhoneBook(names []string, numbers []string) map[string]string {
	cph := make(map[string]string)
	for i, key := range names {
		cph[key] = numbers[i]
	}
	return cph
}

func main() {
	fmt.Println("How many contacts do you want to enter?")
	var a int
	fmt.Scanln(&a)
	for i := 0; i <= a; i++ {
		var names, numbers []string
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter name #%v: ", i+1)
		scanner.Scan()
		s := scanner.Text()
		names = append(names, s)
		fmt.Printf("Enter phone number #%v: ", i+1)
		scanner.Scan()
		s1 := scanner.Text()
		numbers = append(numbers, s1)
		fmt.Println(PhoneBook(names, numbers))
	}
}
