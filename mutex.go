package main

/*import (
	"fmt"
	"sync"
)

/* A Mutex is a method used as a locking mechanism to ensure that only one Goroutine is
accessing the critical section of code at any point of time. This is done to prevent
race conditions from happening. Sync package contains the Mutex. Two methods defined on Mutex
Lock and unlock*/

/*var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}*/
