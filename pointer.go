package main

import "fmt"

func mystructue() {
	m1 := 3
	ptr := &m1
	fmt.Println(ptr)
	fmt.Println(&ptr)
	fmt.Println(*ptr)

}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp

}

/*func main() {
mystructue()
x := 9
y := 6
println(x, y)
swap(&x, &y)
println(x, y)*/

//}
