package main

import (
	"encoding/json"
	"testing"
	// "fmt"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/zerosnake0/jzon"
)

var jsonSmallBytes []byte
var jsonBigBytes []byte
var jsonSmallStruct ABCR
var jsonBigStruct ABCR
var jsoni = jsoniter.ConfigCompatibleWithStandardLibrary

func BenchmarkJsonSetup(b *testing.B) {
	//Create test-data here
	jsonSmallBytes = []byte("{\"a\":\"just\",\"b\":\"another\",\"c\":\"test\"}")
	_ = json.Unmarshal(jsonSmallBytes, &jsonSmallStruct)
	first := ABCR{}
	cur := &first
	for i := 1; i < 1000; i++ {
		sNew := ABCR{
			A: RandStringBytes(3),
			B: RandStringBytes(3),
			C: RandStringBytes(3),
		}
		cur.R = append(cur.R, sNew)
		// cur = &sNew
	}
	jsonBigBytes, _ = json.Marshal(first)
	jsonBigStruct = first

	b.SkipNow()
}

func BenchmarkJsonDefMarshalBig(b *testing.B) {
	for b.Loop() {
		_, _ = json.Marshal(jsonBigStruct)
	}
}

func BenchmarkJsonJsoniterMarshalBig(b *testing.B) {
	for b.Loop() {
		_, _ = jsoni.Marshal(jsonBigStruct)
	}
}

func BenchmarkJsonJzonMarshalBig(b *testing.B) {
	for b.Loop() {
		_, _ = jzon.Marshal(jsonBigStruct)
	}
}

func BenchmarkJsonGojsonMarshalBig(b *testing.B) {
	for b.Loop() {
		_, _ = gojson.Marshal(jsonBigStruct)
	}
}

func BenchmarkJsonDefUnmarshalBig(b *testing.B) {
	var abcr ABCR
	for b.Loop() {
		_ = json.Unmarshal(jsonBigBytes, &abcr)
	}
}

func BenchmarkJsonJsoniterUnmarshalBig(b *testing.B) {
	var abcr ABCR
	for b.Loop() {
		_ = jsoni.Unmarshal(jsonBigBytes, &abcr)
	}
}

func BenchmarkJsonJzonUnmarshalBig(b *testing.B) {
	var abcr ABCR
	for b.Loop() {
		_ = jzon.Unmarshal(jsonBigBytes, &abcr)
	}
}

func BenchmarkJsonGojsonUnmarshalBig(b *testing.B) {
	var abcr ABCR
	for b.Loop() {
		_ = gojson.Unmarshal(jsonBigBytes, &abcr)
	}
}

func BenchmarkJsonDefMarshalSmall(b *testing.B) {
	for b.Loop() {
		_, _ = json.Marshal(jsonSmallStruct)
	}
}

func BenchmarkJsonJsoniterMarshalSmall(b *testing.B) {
	for b.Loop() {
		_, _ = jsoni.Marshal(jsonSmallStruct)
	}
}

func BenchmarkJsonJzonMarshalSmall(b *testing.B) {
	for b.Loop() {
		_, _ = jzon.Marshal(jsonSmallStruct)
	}
}

func BenchmarkJsonGojsonMarshalSmall(b *testing.B) {
	for b.Loop() {
		_, _ = gojson.Marshal(jsonSmallStruct)
	}
}

func BenchmarkJsonDefUnmarshalSmall(b *testing.B) {
	var abcr ABCR
	for b.Loop() {
		_ = json.Unmarshal(jsonSmallBytes, &abcr)
	}
}

func BenchmarkJsonJsoniterUnmarshalSmall(b *testing.B) {
	var abcr ABCR
	for b.Loop() {
		_ = jsoni.Unmarshal(jsonSmallBytes, &abcr)
	}
}

func BenchmarkJsonJzonUnmarshalSmall(b *testing.B) {
	var abcr ABCR
	for b.Loop() {
		_ = jzon.Unmarshal(jsonSmallBytes, &abcr)
	}
}

func BenchmarkJsonGojsonUnmarshalSmall(b *testing.B) {
	var abcr ABCR
	for b.Loop() {
		_ = gojson.Unmarshal(jsonSmallBytes, &abcr)
	}
}
