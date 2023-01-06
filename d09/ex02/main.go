package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func multiplex(in ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	for _, c := range in {
		wg.Add(1)
		c := c
		go func() {
			for d := range c {
				out <- d
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {

	ch1 := make(chan any, 3)
	ch1 <- 111
	ch1 <- "revrfv"
	ch1 <- 4.44
	close(ch1)

	ch2 := make(chan interface{}, 3)
	ch2 <- 222
	ch2 <- "ddssd"
	ch2 <- 5.555
	close(ch2)
	res := multiplex(ch1, ch2)
	for {
		d, ok := <-res
		if !ok {
			break
		}
		fmt.Println(d)
	}

}
