package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign kv
	// emails["jithmi"] = "jithmi@gmail.com"
	// emails["shihara"] = "shihara@gmail.com"
	// emails["jithmishihara"] = "jithmishihara@gmail.com"

	// Decalre map and Add kv
	emails := map[string]string{"jithmi": "jithmi@gamil.com", "shihara": "shihara@gmail.com"}

	emails["jithmishihara"] = "jithmishihara@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["jithmi"])

	// DElete from map
	delete(emails, "jithmi")
	fmt.Println(emails)
}
