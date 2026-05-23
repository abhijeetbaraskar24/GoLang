package main

import "fmt"

func runQ5() {

	var number int

	fmt.Println("enter number : ")
	fmt.Scan(&number)
	count := 0

	if number == 0 {
		count = 1
	} else {
		for number > 0 {
			number = number / 10
			count++
		}
	}

	fmt.Println("total digits are : ", count)

}
