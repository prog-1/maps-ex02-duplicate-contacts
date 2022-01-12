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
	fmt.Print("How many contacts do you want to enter?")
	var cont int
	fmt.Scan(&cont)
	var names, numbers []string
	var snum, sname string
	for st := 1; st <= cont; st++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Enter name #%v: ", st)
		scanner.Scan()
		sname = scanner.Text()
		names = append(names, sname)
		fmt.Printf("Enter phone number #%v: ", st)
		scanner.Scan()
		snum = scanner.Text()
		numbers = append(numbers, snum)
	}
	fmt.Println(createPhoneBook(names, numbers))
}

//func main() {
//	fmt.Print("How many contacts do you want to enter?")
//	var cont int
//	fmt.Scan(&cont)
//	var names, numbers []string
//	for st := 1; st <= cont; st++ {
//		var str string
//		fmt.Printf("Enter name #%v:", st)
//		fmt.Scan(&str)
//		names = append(names, str)
//		fmt.Printf("Enter phone number #%v:", st)
//		fmt.Scan(&str)
//		numbers = append(numbers, str)
//	}
//	fmt.Println(createPhoneBook(names, numbers))
//}
