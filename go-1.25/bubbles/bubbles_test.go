package bubbles_test

import (
	"errors"
	"testing"
	"testing/synctest"
	"time"
)

// BEGIN BUBBLES OMIT
func read(ch <-chan int) (int, error) {
	select {
	case n := <-ch:
		return n, nil
	case <-time.After(60 * time.Second):
		return 0, errors.New("timeout")
	}
}

func TestBubbles(t *testing.T) {
	synctest.Test(t, func(t *testing.T) { // HL
		if _, err := read(make(chan int)); err == nil {
			t.Fatal("expected error")
		}
	})
}

// END BUBBLES OMIT
