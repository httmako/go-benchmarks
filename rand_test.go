package main

import (
	randv1 "math/rand"
	randv2 "math/rand/v2"
	"testing"
	"time"
)

func BenchmarkRandMathV1intn(b *testing.B) {
	for b.Loop() {
		_ = randv1.Intn(56)
	}
}
func BenchmarkRandMathV2intn(b *testing.B) {
	for b.Loop() {
		_ = randv2.IntN(56)
	}
}

func BenchmarkRandMathV1int63(b *testing.B) {
	for b.Loop() {
		_ = randv1.Int63() % 56
	}
}

func BenchmarkRandMathV2int63(b *testing.B) {
	for b.Loop() {
		_ = randv2.Int64() % 56
	}
}

func BenchmarkRandMathV1sourced(b *testing.B) {
	rand := randv1.NewSource(time.Now().UnixNano())
	for b.Loop() {
		_ = rand.Int63() % 56
	}
}

func BenchmarkRandMod(b *testing.B) {
	mod := randv1.Int() % 112
	for mod%2 == 0 || mod%7 == 0 {
		mod += 1
	}
	counter := 0
	for b.Loop() {
		counter = (counter + mod) % 56
	}
}
