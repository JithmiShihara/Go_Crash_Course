package main

import "fmt"

func main() {

	//var name = "Jithmi"
	var age = 22
	const isCool = true
	var size float32 = 2.3

	name, email := "Jithmi", "jithmi@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
