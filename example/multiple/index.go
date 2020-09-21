package main

import (
	"time"

	. "github.com/nmccready/go-debug"
	"github.com/nmccready/go-debug/example/rootDebug"
)

func work(debug *Debugger, delay time.Duration) {
	for {
		debug.Log("doing stuff")
		debug.Log("<body>Hi</body>")
		debug.Log(Fields{
			"oh": "hi",
		})
		debug.Error("oops")
		time.Sleep(delay)
	}
}

func main() {
	a := rootDebug.Spawn("multiple:a").WithFields(Fields{
		"junk":    "hi junk",
		"another": 1,
	})
	var b = rootDebug.Spawn("multiple:b")
	var c = rootDebug.Spawn("multiple:c")

	q := make(chan bool)

	go work(a, 1000*time.Millisecond)
	go work(b, 250*time.Millisecond)
	go work(c, 100*time.Millisecond)

	<-q
}
