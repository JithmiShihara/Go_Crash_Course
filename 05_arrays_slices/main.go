package main

import "fmt"

func main() {
	// Array
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])

	// Declare and Assign
	flower := [2]string{"Rose", "Lily"}

	fmt.Println(flower)

	animalSlice := []string{"cat", "Dog", "Tiger", "elephant"}

	fmt.Println(len(animalSlice))
	fmt.Println(animalSlice[1:2])

	fmt.Println(animalSlice[0:3])

}
