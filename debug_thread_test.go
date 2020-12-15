package debug

import (
	"bytes"
	"strconv"
	"sync"
	"testing"
)

func TestThreadSafety(t *testing.T) {
	debug := Debug("foo")
	var b []byte
	buf := bytes.NewBuffer(b)
	SetWriter(buf)
	Enable("*")

	wg := sync.WaitGroup{}
	totalThreads := 8000
	wg.Add(totalThreads)

	for i := 0; i < totalThreads; i++ {
		i := i
		go func() {
			defer wg.Done()
			f := Fields{}
			f[strconv.Itoa(i)] = "test"
			debug.Spawn("thread" + strconv.Itoa(i)).WithFields(f).Log("something")
		}()
	}
	wg.Wait()
}

func TestThreadJsonSafety(t *testing.T) {
	debug := Debug("foo")
	SetFormatter(&JSONFormatter{})
	var b []byte
	buf := bytes.NewBuffer(b)
	SetWriter(buf)
	Enable("*")

	wg := sync.WaitGroup{}
	totalThreads := 1000
	wg.Add(totalThreads)

	for i := 0; i < totalThreads; i++ {
		i := i
		go func() {
			defer wg.Done()
			f := Fields{}
			f[strconv.Itoa(i)] = "test"
			f["a"] = "test"
			f["b"] = "test"
			f["c"] = "test"
			f["d"] = "test"
			debug.Spawn("thread" + strconv.Itoa(i)).WithFields(f).Log("something")
		}()
	}
	wg.Wait()
}
