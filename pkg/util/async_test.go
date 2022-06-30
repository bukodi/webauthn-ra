package util

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch := make(chan string)
	//go PrintStringCh(ch)
	go PrintCh(ch)
	ch <- "first"
	ch <- "second"
	ch <- "third"
	close(ch)
	time.Sleep(time.Millisecond * 10)
}

func PrintStringCh(ch <-chan string) {

loop:
	for {
		select {
		case s, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed.")
				break loop
			}
			fmt.Println(s)
		}
	}
}

func PrintCh[T any](ch <-chan T) {

loop:
	for {
		select {
		case s, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed.")
				break loop
			}
			fmt.Println(s)
		}
	}
}
