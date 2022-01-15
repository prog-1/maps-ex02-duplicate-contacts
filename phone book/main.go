package main

import (
	"bufio"
	"fmt"
	"os"
)

func createPhoneBook(names []string, numbers []string) map[string]string {
	x := make(map[string]string)
	y := len(names)
	for i := 0; i < y; i++ {
		x[names[i]] = numbers[i]
	}
	return x
}
func main() {
	fmt.Println("How many contacts do you want to enter")
	var cont int
	fmt.Scanln(&cont)
	var g string
	var names, numbers []string
	for i := 0; i < cont; i++ {
		fmt.Printf("Enter name:", i+1)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		g = scanner.Text()
		names = append(names, g)
		g = scanner.Text()
		names = append(names, g)
		scanner.Scan()
		g = scanner.Text()
		fmt.Printf("Enter phone:", i+1)
		numbers = append(numbers, g)
	}
	fmt.Println(createPhoneBook(numbers, names))
}
