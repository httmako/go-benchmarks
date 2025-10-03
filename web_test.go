package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var nethttp *http.ServeMux
var r *mux.Router
var app *gin.Engine
var e *echo.Echo
var byteReader []byte

func TestMain(m *testing.M) {
	// net/http
	nethttp = http.NewServeMux()
	nethttp.HandleFunc("GET /fmt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})
	nethttp.HandleFunc("GET /w", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	nethttp.HandleFunc("GET /abc", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	})
	nethttp.HandleFunc("POST /json", func(w http.ResponseWriter, r *http.Request) {
		var abc ABC
		err := json.NewDecoder(r.Body).Decode(&abc)
		if err != nil {
			panic(err)
		}
	})
	// gin
	gin.SetMode(gin.ReleaseMode)
	app = gin.New()
	app.GET("/abc", func(c *gin.Context) {
		c.String(200, "OK")
	})
	app.POST("/json", func(c *gin.Context) {
		var abc ABC
		err := c.BindJSON(&abc)
		if err != nil {
			panic(err)
		}
	})
	// echo
	e = echo.New()
	e.GET("/abc", func(c echo.Context) error {
		return c.String(200, "OK")
	})
	e.POST("/json", func(c echo.Context) error {
		var abc ABC
		err := c.Bind(&abc)
		if err != nil {
			panic(err)
		}
		return nil
	})
	r = mux.NewRouter()
	r.HandleFunc("/abc", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	})

	byteReader = []byte("{\"a\":\"just\",\"b\":\"another\",\"c\":\"test\"}")

	m.Run()
}

func BenchmarkNetHttpGetFMT(b *testing.B) {
	r := httptest.NewRequest("GET", "http://example.com/fmt", nil)
	w := httptest.NewRecorder()
	for b.Loop() {
		nethttp.ServeHTTP(w, r)
	}
}

func BenchmarkNetHttpGetIO(b *testing.B) {
	r := httptest.NewRequest("GET", "http://example.com/abc", nil)
	w := httptest.NewRecorder()
	for b.Loop() {
		nethttp.ServeHTTP(w, r)
	}
}

func BenchmarkNetHttpGetWrite(b *testing.B) {
	r := httptest.NewRequest("GET", "http://example.com/w", nil)
	w := httptest.NewRecorder()
	for b.Loop() {
		nethttp.ServeHTTP(w, r)
	}
}

func BenchmarkGinGet(b *testing.B) {
	r := httptest.NewRequest("GET", "http://example.com/abc", nil)
	w := httptest.NewRecorder()
	for b.Loop() {
		app.ServeHTTP(w, r)
	}
}

func BenchmarkEchoGet(b *testing.B) {
	r := httptest.NewRequest("GET", "http://example.com/abc", nil)
	w := httptest.NewRecorder()
	for b.Loop() {
		e.ServeHTTP(w, r)
	}
}

func BenchmarkGorillaMuxGet(b *testing.B) {
	req := httptest.NewRequest("GET", "http://example.com/abc", nil)
	w := httptest.NewRecorder()
	for b.Loop() {
		r.ServeHTTP(w, req)
	}
}

// JSON POST
// The Content-Type must be set because echo panic's without it (gin/default are fine)
func BenchmarkDefPostJson(b *testing.B) {
	for b.Loop() {
		r := httptest.NewRequest("POST", "http://example.com/json", bytes.NewReader(byteReader))
		r.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()
		nethttp.ServeHTTP(w, r)
	}
}

func BenchmarkGinPostJson(b *testing.B) {
	for b.Loop() {
		r := httptest.NewRequest("POST", "http://example.com/json", bytes.NewReader(byteReader))
		r.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()
		app.ServeHTTP(w, r)
	}
}

func BenchmarkEchoPostJson(b *testing.B) {
	for b.Loop() {
		r := httptest.NewRequest("POST", "http://example.com/json", bytes.NewReader(byteReader))
		r.Header.Add("Content-Type", "application/json")
		w := httptest.NewRecorder()
		e.ServeHTTP(w, r)
	}
}
