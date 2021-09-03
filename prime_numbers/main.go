package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func testForPrime(c int, ps chan int, done chan bool) {
	var nps chan int
	var stopper chan bool
	fmt.Printf("%d is prime\n", c)
	defer wg.Done()
	for true {
		select {
		case <-done:
			fmt.Println("Closing down")
			if stopper != nil {
				stopper <- true
			}
			break
		default:
			p := <-ps
			if p%c == 0 {
				continue
			} else if nps == nil {
				nps = make(chan int)
				stopper = make(chan bool)
				wg.Add(1)
				go testForPrime(p, nps, stopper)
			} else {
				nps <- p
			}
		}
	}
}

func main() {
	fmt.Println("Lets go")
	possible_primes := make(chan int)
	stopper := make(chan bool)
	wg.Add(1)
	go testForPrime(2, possible_primes, stopper)
	for i := 3; i < 30; i++ {
		possible_primes <- i
	}
	stopper <- true
	wg.Wait()
}
