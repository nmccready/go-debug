package integration

import (
	"bytes"
	"io"
	"strconv"
	"sync"
	"testing"

	. "github.com/nmccready/go-debug"
)

func TestThreadSafety(t *testing.T) {
	debug := Debug("foo")
	Enable("*")

	var w io.Writer = bytes.NewBuffer([]byte{})
	writers := []*io.Writer{nil, &w} // test both buffer and stdout

	for _, writerPtr := range writers {
		writerPtr := writerPtr
		if writerPtr != nil {
			SetWriter(*writerPtr)
		}

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
}

func TestThreadJsonSafety(t *testing.T) {
	debug := Debug("foo")
	SetFormatter(&JSONFormatter{})
	Enable("*")

	var w io.Writer = bytes.NewBuffer([]byte{})
	writers := []*io.Writer{nil, &w} // test both buffer and stdout

	for _, writerPtr := range writers {
		writerPtr := writerPtr
		if writerPtr != nil {
			SetWriter(*writerPtr)
		}

		wg := sync.WaitGroup{}
		totalThreads := 7000
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

}
