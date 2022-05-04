package main

import (
	"fmt"
	"log"
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
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var numbers, names []string
	var name, number string
	for {
		_, err := fmt.Fscan(f, &name, &number)
		if err != nil {
			break
		}
		names = append(names, name)
		numbers = append(numbers, number)
	}
	fmt.Println(createPhoneBook(names, numbers))
}
