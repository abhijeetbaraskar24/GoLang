package main

import "fmt"

func main() {
	var choice int
	var x int
	var y int

	fmt.Println("1.Addition\n2.Substraction\n3.Multiplication\n4.Division\n5.Square\n6.Exit")

	

	for {

		fmt.Println("Enter your choice : ")
		fmt.Scan(&choice)


		if choice == 1{
			fmt.Println("enter first number : ")
			fmt.Scan(&x)
			fmt.Println("enter second number : ")
			fmt.Scan(&y)
			add(x,y)
		}else if choice == 2{
			fmt.Println("enter first number : ")
			fmt.Scan(&x)
			fmt.Println("enter second number : ")
			fmt.Scan(&y)
			sub(x,y)
		}else if choice == 3{
			fmt.Println("enter first number : ")
			fmt.Scan(&x)
			fmt.Println("enter second number : ")
			fmt.Scan(&y)
			mult(x,y)
		}else if choice == 4{
			fmt.Println("enter first number : ")
			fmt.Scan(&x)
			fmt.Println("enter second number : ")
			fmt.Scan(&y)
			div(x,y)
		}else if choice == 5 {
			fmt.Println("enter number : ")
			fmt.Scan(&x)
			square(x)
		} else{
			fmt.Println("Program End ...")
			break 
		}
	}
	
}

func add(a int, b int){
	result := a + b
	fmt.Println("Addition = ",result)
}

func sub(a int, b int){
	result := a - b
	fmt.Println("Subtraction = ",result)
}

func mult(a int, b int){
	result := a * b
	fmt.Println("Multiplication = ",result)
}

func div(a int, b int ){
	if b == 0 {
		fmt.Println("Cannot divide by zero")
		return
	}

	result := a / b
	fmt.Println("Division = ",result)
}

func square(a int){
	result := a * a
	fmt.Println("Square = ",result)
}
