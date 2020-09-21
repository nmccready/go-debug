package main

import (
	"time"

	. "github.com/nmccready/go-debug"
	"github.com/nmccready/go-debug/example/rootDebug"
)

func work(debug *Debugger, delay time.Duration) {
	for {
		debug.Log("a string")
		// debug.Log(Fields{
		// 	"one":   1,
		// 	"bool":  true,
		// 	"junk1": "hi junk1",
		// })
		time.Sleep(delay)
	}
}

func main() {
	SetFormatter(&JSONFormatter{PrettyPrint: true, FlattenMsgFields: true})
	a := rootDebug.Spawn("multiple:a").WithFields(Fields{
		"global": "global field",
	})
	// fmt.Printf("fields %s\n", a.)
	var b = rootDebug.Spawn("multiple:b")
	var c = rootDebug.Spawn("multiple:c")

	q := make(chan bool)

	go work(a, 1000*time.Millisecond)
	go work(b, 250*time.Millisecond)
	go work(c, 100*time.Millisecond)

	<-q
}
