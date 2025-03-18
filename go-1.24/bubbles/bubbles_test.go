package bubbles

import (
	"fmt"
	"testing"
	"testing/synctest"
	"time"
)

// BEGIN RETRY OMIT
func Retry(f func()) {
	done := time.After(30 * time.Second)

	for {
		f()

		select {
		case <-time.After(10 * time.Second):
			continue
		case <-done:
			return
		}
	}
}

// END RETRY OMIT

func checkEquals(t *testing.T, got, expected int) {
	if got != expected {
		t.Fatalf("got %d, want %d", got, expected)
	}
}

func TestRetry(t *testing.T) {
	var times int
	// BEGIN RETRY_TEST OMIT
	synctest.Run(func() {
		go Retry(func() { times += 1 })
		synctest.Wait()
		checkEquals(t, times, 1) // f() is called immediately.
		time.Sleep(5 * time.Second)
		synctest.Wait()
		checkEquals(t, times, 1) // 5s later: the goroutine is still waiting.
		time.Sleep(25 * time.Second)
		synctest.Wait()
		checkEquals(t, times, 4) // 25s later: f() is called 3x more.
		time.Sleep(10 * time.Second)
		synctest.Wait()
		checkEquals(t, times, 4) // 10s later: f() is not called.
	})
	// END RETRY_TEST OMIT
}

// BEGIN TIMED_READ OMIT
func TimedRead(ch <-chan int) (int, error) {
	select {
	case <-time.After(10 * time.Second):
		return 0, fmt.Errorf("expired")
	case n := <-ch:
		return n, nil
	}
}

// END TIMED_READ OMIT

func checkNil(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func checkNotNil(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("no error returned")
	}
}

func TestTimedRead(t *testing.T) {
	var n int
	var err error

	// BEGIN TEST_TIMED_READ OMIT
	synctest.Run(func() {
		// Reading nil channels always blocks.
		go func() { n, err = TimedRead(nil) }()
		// 5s later: still waiting...
		time.Sleep(5 * time.Second)
		synctest.Wait()
		checkNil(t, err)
		// 5s later: channel returned nothing, error!
		time.Sleep(5 * time.Second)
		synctest.Wait()
		checkNotNil(t, err)
	})
	// END TEST_TIMED_READ OMIT

	_ = n
}
