package main

import "fmt"

func runQ3() {

	var num int
	total := 0
	fmt.Print("enter num : ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {

		total = total + i

	}

	fmt.Print("total is : ", total)

}
