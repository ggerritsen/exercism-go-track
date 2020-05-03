// Package letter provides functionality around letter frequencies in texts
package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

type concFreqMap struct {
	sync.Mutex
	internal FreqMap
}

func (m *concFreqMap) add(r rune, i int) {
	m.Lock()
	defer m.Unlock()
	m.internal[r] += i
}

// ConcurrentFrequency counts the frequency of each rune in a given text concurrently and returns this
// data as a FreqMap.
func ConcurrentFrequency(s []string) FreqMap {
	m := concFreqMap{internal: FreqMap{}}

	var wg sync.WaitGroup
	for _, text := range s {
		wg.Add(1)
		t := text
		go func() {
			f := Frequency(t)
			for k, v := range f {
				m.add(k, v)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return m.internal
}
