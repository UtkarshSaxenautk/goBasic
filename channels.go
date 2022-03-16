package main

/*import (
	"fmt"
	"time"
)*/

// Unbuffered channel
/*func main() {
	// <name> chan <datatype>
	channel := make(chan int)
	var input int
	fmt.Scanln(&input)
	go func() {

		channel <- input
	}()
	val := <-channel
	fmt.Println(val)
	fmt.Println(channel)
	time.Sleep(time.Second * 5)
	go func() {

		fmt.Scanln(&input)
		channel <- input
	}()
	val = <-channel
	fmt.Println(val)
	fmt.Println(channel)
}*/

/*type car struct {
	Model  string
	Color  string
	Number int
}

func main() {
	channel := make(chan *car, 3)

	go func() {
		channel <- &car{Model: "BMW", Color: "Red", Number: 50}
		channel <- &car{Model: "Audi", Color: "Blue", Number: 100}
		channel <- &car{Model: "Porsche", Color: "Black", Number: 150}
		channel <- &car{Model: "Lamborgini", Color: "White", Number: 200}
		close(channel)
	}()
	for i := range channel {
		fmt.Println("Car is of number : ", i.Number, " and is of Model name:  ", i.Model, " with color: ", i.Color)
	}
}*/
