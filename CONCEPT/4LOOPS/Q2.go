package main

import "fmt"

func runQ2() {

	var n int
	fmt.Print("enter num : ")
	fmt.Scan(&n)

	x := 1

	for i := x; i < 11; i++ {
		fmt.Println(n, "X", i, "=", n*i)
	}

}
