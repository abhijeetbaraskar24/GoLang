package main

import "fmt"

func runQ6() {
	var snum int
	fmt.Print("enter number : ")
	fmt.Scan(&snum)

	digit := 0
	sum := 0

	if snum == 0 {
		sum = 0
	} else {
		for snum > 0 {
			digit = snum % 10
			sum = sum + digit
			snum = snum / 10
		}
	}

	fmt.Print("total sum is ", sum)
}
