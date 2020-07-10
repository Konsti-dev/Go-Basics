package main

import "fmt"

func main() {
	// Define map

	/* email := make(map[string]string)

	// Assign kv

	email["Bob"] = "bob@gmail.com"
	email["Alice"] = "Alice@gmail.com"
	email["Mike"] = "Mike@gmail.com" */

	// Declare map and add kv

	email := map[string]string{"Bob": "bob@gmail.com", "Alice": "Alice@gmail.com"}

	email["Mike"] = "Mike@gmail.com"

	fmt.Println(email)
	fmt.Println(len(email))
	fmt.Println(email["Bob"])

	// Delete from map
	delete(email, "Bob")
	fmt.Println(email)
}
