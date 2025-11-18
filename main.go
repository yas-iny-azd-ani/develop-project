package main

import "fmt"

func main() {
	fmt.Println("Hello")

	var name string
	fmt.Print("enter your name: ")
	fmt.Scan(&name)

	var age int
	fmt.Print("enter your age: ")
	fmt.Scan(&age)

	var height float64
	fmt.Print("enter your height: ")
	fmt.Scan(&height)
}
