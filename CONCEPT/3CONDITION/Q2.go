package main

import "fmt"

func runQ2() {

	var marks int

	fmt.Print("enter marks : ")
	fmt.Scan(&marks)

	if marks > 90 {
		fmt.Print("A +")
	} else if marks < 91 && marks > 75 {
		fmt.Print("B")
	} else if marks < 76 && marks > 50 {
		fmt.Print("C")
	} else {
		fmt.Print("fail")
	}

}
