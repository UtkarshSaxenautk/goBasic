package main

/*import (
	"fmt"
	"os"
	"time"
)

func selector(ch chan int, st chan struct{}) {
	time.Sleep(time.Second)
	for {
		select {
		case <-ch:
			fmt.Println("Got a int")
		case <-st:
			fmt.Println("Got a structure")
			os.Exit(0)
		}
	}
}

type student struct {
	Rollno int
	Name   string
}

func main() {
	channel_int := make(chan int, 2)
	channel_struct := make(chan struct{})

	go func() {
		channel_int <- 1
		channel_int <- 2
		channel_struct <- struct{}{}

	}()
	go selector(channel_int, channel_struct)
	select {}
}*/
