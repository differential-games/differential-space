package game

import (
	"math/rand"
	"testing"
	"time"
)

func TestOptions(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	fails := 0
	for i := 0; i < 100; i++ {
		_, err := New(DefaultOptions)
		if err != nil {
			t.Error(err)
			fails++
		}
	}
	if fails > 0 {
		t.Errorf("Got %d%% failure rate", fails)
	}
}
