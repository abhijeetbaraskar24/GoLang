package main

import "fmt"

func runQ1() {
	var age int

	fmt.Println("enter ur age : ")
	fmt.Scan(&age)

	if age > 17 {
		fmt.Println("adult")
	} else if age < 18 && age > 12 {
		fmt.Println("teen")
	} else {
		fmt.Println("child")
	}

}
