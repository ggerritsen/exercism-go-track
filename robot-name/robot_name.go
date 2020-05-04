// Package robotname provides functionality around creating random and unique robot names
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		n, err :=  genName()
		if err != nil {
			return "", err
		}
		r.name = n
	}

	return r.name, nil
}

var alphabet = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
}

var existing = map[string]bool{}

var max int = 26*26*10*10*10

func genName() (string, error) {
	if len(existing) >= max {
		return "", fmt.Errorf("exhausted")
	}

	name := fmt.Sprintf("%s%s%d%d%d",
		alphabet[rand.Intn(26)],
		alphabet[rand.Intn(26)],
		rand.Intn(10),
		rand.Intn(10),
		rand.Intn(10))

	if existing[name] {
		return genName()
	}

	existing[name] = true
	return name, nil
}

// Reset resets a robot to factory settings, also creating a new name
func (r *Robot) Reset() {
	r.name = ""
}
