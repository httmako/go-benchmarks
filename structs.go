package main

import (
	"math/rand"
)

type ABC struct {
	A string `json:"a" query:"a"`
	B string `json:"b" query:"b"`
	C string `json:"c" query:"c"`
}

type ABCR struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	R []ABCR `json:"abcr"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
