package main

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

var wpmux *http.ServeMux

func BenchmarkPrepareWebpanic(b *testing.B) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	wpmux = http.NewServeMux()
	wpmux.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			re := recover()
			logger.Info("webreq", "url", r.URL, "err", re)
		}()
		http.Error(w, "server error", 500)
		panic("ay")
	})
	wpmux.HandleFunc("GET /log", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			re := recover()
			logger.Info("webreq", "url", r.URL, "err", re)
		}()
		http.Error(w, "server error", 500)
		logger.Error("error", "err", "ay")
	})
	b.SkipNow()
}

func BenchmarkWebPanic(b *testing.B) {
	r := httptest.NewRequest("GET", "http://example.com/panic", nil)
	w := httptest.NewRecorder()
	for b.Loop() {
		wpmux.ServeHTTP(w, r)
	}
}

func BenchmarkWebLog(b *testing.B) {
	r := httptest.NewRequest("GET", "http://example.com/log", nil)
	w := httptest.NewRecorder()
	for b.Loop() {
		wpmux.ServeHTTP(w, r)
	}
}
