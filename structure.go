package main

import "fmt"

type car struct {
	Name     string
	age      int
	model_no int
}

func (c car) Info() {
	fmt.Println("I am driving ", c.Name, " of model ", c.model_no, " with age : ", c.age)
}

func (c car) Getname() string {
	return c.Name
}

func (c car) Getage() int {
	return c.age
}

/*func main() {
	c := car{
		Name:     "BMW",
		age:      4,
		model_no: 3409,
	}
	fmt.Println(c)
	c.Info()
	name_of_car := c.Getname()
	age_of_car := c.Getage()
	fmt.Println(name_of_car)
	fmt.Println(age_of_car)
}*/
