package main

import (
	"time"

	. "github.com/nmccready/go-debug"
	"github.com/nmccready/go-debug/example/rootDebug"
)

func work(debug *Debugger, delay time.Duration) {
	for {
		debug.Log("doing stuff")
		time.Sleep(delay)
	}
}

func main() {
	SetFormatter(&JSONFormatter{})
	a := rootDebug.Spawn("multiple:a").WithFields(Fields{
		"junk":     "hi junk",
		"another":  1,
		"another2": 2,
		"junk1":    "hi junk1",
	})
	// fmt.Printf("fields %s\n", a.)
	var b = rootDebug.Spawn("multiple:b")
	var c = rootDebug.Spawn("multiple:c")
	var d = rootDebug.Spawn("multiple:d")

	q := make(chan bool)

	go work(a, 500*time.Millisecond)
	go work(b, 250*time.Millisecond)
	go work(c, 100*time.Millisecond)
	go work(d, 120*time.Millisecond)

	<-q
}
