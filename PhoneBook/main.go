package main

import "fmt"

func createPhoneBook(names []string, numbers []string) map[string]string {
	for i := 0,i <= len(names) && len(numbers),i++{
		map += names[i]string + numbers[i]string 
	}
	for indx := 0,indx <= len(names),indx++{
		for indxx := 0,indxx <= len (numbers),indxx++{
			fmt.Println("Something wrong. You entered same numbers or names")
		}
	}
	return map
}

func main() {
	var con,r int
	name := make([]string)
	numbers := make([]string)
	fmt.Println("How many contacts do you want to enter?")
	fmt.Scan(&con)
	for r == 1,r < con+1,r++{
		fmt.Println("Enter name #",r,":")
		name += reader.ReadString('\n')
		fmt.Println("Enter phone number #",r,":")
		numbers += reader.ReadString('\n')
	}
	return name,numbers
	fmt.Println("Map: ",createPhoneBook(map)) 
}
