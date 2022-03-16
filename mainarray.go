package main

import "fmt"

// array
func array() {
	a := []int{2, 8, 6}
	b := []string{"Hello", "World", "and", "Java"}
	b = append(b, "This", "is", "my", "fav", "language")
	fmt.Println(a, b)
}

/*func main() {
	array()
}*/
