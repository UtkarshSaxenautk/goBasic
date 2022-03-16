package main

/*import (
	"fmt"
	"time"
)

func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("Ping")
		time.Sleep(time.Second * 1)
		ponger <- 1
	}
}

func ponger(ponger <-chan int, pinger chan<- int) {
	for {
		<-ponger
		fmt.Println("Pong")
		time.Sleep(time.Second * 1)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(pong, ping)

	ping <- 1
	// we can use this or select routine
	/*for {
		time.Sleep(time.Second)
	}*/
//	select {}
//}
