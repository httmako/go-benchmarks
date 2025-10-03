package main

import (
	"testing"
)

func BenchmarkChannels(b *testing.B) {
	c := make(chan int)
	go func() {
		for {
			<-c
		}
	}()
	for b.Loop() {
		c <- 1
	}
}
func BenchmarkChannelsBuffered(b *testing.B) {
	cb := make(chan int, 100)
	go func() {
		for {
			<-cb
		}
	}()
	for b.Loop() {
		cb <- 1
	}
}
