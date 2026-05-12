package main

import "fmt"

func runQ4() {
	var no int
	fmt.Print("enter num : ")
	fmt.Scan(&no)

	rev := 0
	digit := 0

	for no > 0 {
		digit = no % 10
		rev = rev*10 + digit

		no = no / 10
	}
	fmt.Print(rev)

}
