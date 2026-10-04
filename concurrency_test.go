package main

import (
	"sync/atomic"
	"testing"
)

/*
This test is supposed to compare 2 methods of doing goroutine synchronization.
Use case: Multiple goroutines have to do a thing X times alltogether.
1. Use a channel and feed into it only as much as you need to be done
2. Use an atomic Int and set it to the desired count, then count down until you reach 0
*/

func doNothing() {}

func BenchmarkConcurrencyChan(b *testing.B) {
	c := make(chan int, 100)
	go func() {
		for {
			c <- 1
		}
	}()
	for b.Loop() {
		<-c
		doNothing()
	}
}
func BenchmarkConcurrencySync(b *testing.B) {
	c := atomic.Int64{}
	for b.Loop() {
		if n := c.Add(1); n > 0 {
			doNothing()
		}
	}
}
