package main

import (
	"fmt"
	"sort"
)

func find(name string, names []string) int {
	for i, n := range names {
		if n == name {
			return i
		}
	}
	return -1
}

func createPhoneBook(names []string, numbers []string) map[string]string {
	pb := make(map[string]string)
	for i, name := range names {
		if find(name, names) == -1 {
			pb[names[i]] = numbers[i]
		} else {
			fi := find(name, names)
			pb[names[fi]] = numbers[i]
		}
	}
	return pb
}

func main() {
	fmt.Print("How many contacts do you want to enter?")
	var c int
	fmt.Scan(&c)
	var names, numbers []string
	var tmp string
	for i := 1; i <= c; i++ {
		fmt.Printf("Enter name #%v:", i)
		fmt.Scan(&tmp)
		names = append(names, tmp)
		fmt.Printf("Enter phone number #%v:", i)
		fmt.Scan(&tmp)
		numbers = append(numbers, tmp)
	}
	pb := createPhoneBook(names, numbers)
	var keys []string
	for k := range pb {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, ":", pb[k])
	}
}
