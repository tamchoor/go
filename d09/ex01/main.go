package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

func crawlWeb(inputCh <-chan string, ctx context.Context) <-chan *string {
	runtime.GOMAXPROCS(8)
	var wg sync.WaitGroup

	outputCh := make(chan *string, len(inputCh))
	defer close(outputCh)
	for url := range inputCh {
		fmt.Println(url)
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				resp, err := http.Get(url)
				if err != nil {
					log.Fatalln(err)
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
				res := string(body)
				outputCh <- &res
			}
		}(url)
	}
	wg.Wait()
	return outputCh

}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	input := make(chan string, 2)
	input <- "http://localhost:8888"
	input <- "http://golang.org/"
	close(input)

	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("\nMy job here is done!")
		cancel()
	}()

	outCh := crawlWeb(input, ctx)
	for res := range outCh {
		fmt.Println(len(outCh), "->\t", *res)
	}
}
