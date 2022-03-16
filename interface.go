package main

import "fmt"

type vechiles interface {
	drive()
	stop()
}

type s struct {
	nameofcar string
	model_no  int
}

type bike struct {
	nameofbike string
	model_no   int
}

func (c *s) drive() {
	fmt.Println(c.nameofcar, " is on the move ")
	fmt.Println(c.nameofcar, "has model no : ", c.model_no)
}

func (b *bike) drive() {
	fmt.Println(b.nameofbike, " is on move")
	fmt.Println(b.nameofbike, "has model no : ", b.model_no)
}
func (c *s) stop() {
	fmt.Println(c.nameofcar, "is going to stop")
	fmt.Println(c.nameofcar, "has model no : ", c.model_no)
}
func (b *bike) stop() {
	fmt.Println(b.nameofbike, "is going to stop")
	fmt.Println(b.nameofbike, "has model no : ", b.model_no)
}

/*func main() {
	c := car{
		nameofcar: "BMW",
		model_no:  8976,
	}
	b := bike{
		nameofbike: "Ducati",
		model_no:   7865,
	}

	b.drive()
	c.drive()
	b.stop()
	c.stop()
}*/
