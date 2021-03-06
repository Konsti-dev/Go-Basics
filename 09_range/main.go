package main

import "fmt"

func main() {
	ids := []int{33, 55, 66, 77, 7}

	// loop through ids

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	email := map[string]string{"Bob": "bob@gmail.com", "Alice": "Alice@gmail.com"}

	for k, v := range email {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range email {
		fmt.Println("Name: " + k)
	}

}
