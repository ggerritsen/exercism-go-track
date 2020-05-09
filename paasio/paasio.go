// Package paasio provides functionality around counting operations
package paasio

import (
	"io"
	"sync"
)

type myWriteCounter struct {
	sync.RWMutex
	w     io.Writer
	bytes int64
	nops  int
}

func (wc *myWriteCounter) Write(p []byte) (int, error) {
	n, err := wc.w.Write(p)

	wc.Lock()
	defer wc.Unlock()
	wc.bytes += int64(n)
	wc.nops++

	return n, err
}

func (wc *myWriteCounter) WriteCount() (int64, int) {
	wc.RLock()
	defer wc.RUnlock()
	return wc.bytes, wc.nops
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &myWriteCounter{w: w}
}

type myReadCounter struct {
	sync.RWMutex
	r     io.Reader
	bytes int64
	nops  int
}

func (rc *myReadCounter) Read(p []byte) (int, error) {
	n, err := rc.r.Read(p)

	rc.Lock()
	defer rc.Unlock()
	rc.bytes += int64(n)
	rc.nops++

	return n, err
}

func (rc *myReadCounter) ReadCount() (n int64, nops int) {
	rc.RLock()
	defer rc.RUnlock()
	return rc.bytes, rc.nops
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &myReadCounter{r: r}
}

type myReadWriteCounter struct {
	r ReadCounter
	w WriteCounter
}

func (rwc *myReadWriteCounter) Read(p []byte) (int, error) {
	return rwc.r.Read(p)
}

func (rwc *myReadWriteCounter) Write(p []byte) (int, error) {
	return rwc.w.Write(p)
}

func (rwc *myReadWriteCounter) ReadCount() (n int64, nops int) {
	return rwc.r.ReadCount()
}

func (rwc *myReadWriteCounter) WriteCount() (n int64, nops int) {
	return rwc.w.WriteCount()
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &myReadWriteCounter{r: NewReadCounter(rw), w: NewWriteCounter(rw)}
}
