package examples

import (
	"database/sql"
	"iter"
	"maps"
	"slices"
	"time"
	"unique"
)

func useRange() {
	// BEGIN USE RANGE OMIT
	for i, v := range [3]int{1, 2, 3} {
		// i = index, v = value
		_, _ = i, v // OMIT
	}
	for i, x := range []int{1, 2, 3} {
		// i = index, v = value
		_, _ = i, x // OMIT
	}
	for k, v := range map[int]int{1: 2, 3: 4} {
		// k = key, v = value
		_, _ = k, v // OMIT
	}
	for i, r := range "some text" {
		// i = index, r = rune
		_, _ = i, r // OMIT
	}
	for x := range make(chan struct{}) {
		// x = element
		_ = x // OMIT
	}
	for i := range 42 { // Since Go 1.22
		// i = integer in [0, 42)
		_ = i // OMIT
	}
	// END USE RANGE OMIT
}

func useCustomAPI(db *sql.DB) error {
	// BEGIN USE API OMIT
	// Get a cursor to the beginning of the results.
	rows, err := db.Query("SELECT id FROM users")
	if err != nil {
		return err
	}

	// Always close the cursor!
	defer rows.Close()

	// Iterate over the rows.
	for rows.Next() {
		// Use the next row...
	}

	// Check if the iteration failed.
	if err := rows.Err(); err != nil {
		return err
	}
	// END USE API OMIT
	return nil
}

// BEGIN ITER DEF OMIT
type (
	Seq[V any]     func(yield func(V) bool)
	Seq2[K, V any] func(yield func(K, V) bool)
)

func consumeElements[T any](iterator iter.Seq[T]) {
	for i := range iterator {
		// ...
		_ = i // OMIT
	}
}

func consumePairs[K, V any](iterator iter.Seq2[K, V]) {
	for k, v := range iterator {
		// ...
		_, _ = k, v // OMIT
	}
}

// END ITER DEF OMIT

func iteratorExamples() {
	// BEGIN ITER EXAMPLE OMIT
	// Iterator returning one element (iter.Seq)
	for k := range maps.Keys(map[int]int{1: 2, 3: 4}) {
		// k = 1, 2
		_ = k // OMIT
	}

	// Iterator returning two elements (iter.Seq2)
	for i, v := range slices.Backward([]int{1, 2, 3}) {
		// i, v = (2, 3), (1, 2), (0, 1)
		_, _ = i, v // OMIT
	}

	// Iterator returning 2-elements long slices (iter.Seq)
	for chunk := range slices.Chunk([]int{1, 2, 3}, 2) {
		// chunk = {1, 2}, {3}
		_ = chunk // OMIT
	}
	// END ITER EXAMPLE OMIT
}

type Data struct{}

// BEGIN INTERN OMIT
type IPv6 [16]uint8

type Connection struct {
	Origin unique.Handle[IPv6]
}

func newConnection(origin IPv6) Connection {
	// Smaller data types use less memory.
	return Connection{Origin: unique.Make(origin)}
}

func sameOrigin(c1, c2 Connection) bool {
	// Faster comparisons use less CPU.
	return c1.Origin == c2.Origin
}

// END INTERN OMIT

func timerProblems() {
	// BEGIN TIMER GC OMIT
	inOneHour := time.NewTimer(1 * time.Hour)
	inOneMinute := time.NewTimer(1 * time.Minute)
	select {
	case <-inOneHour.C:
		// One hour passed.
	case <-inOneMinute.C:
		// But one minute passes faster.
	}
	// inOneHour will be GC'd... in 59 minutes.
	// END TIMER GC OMIT
}

func tickerProblems() {
	// BEGIN TICKER GC OMIT
	inOneMinute := time.NewTimer(1 * time.Minute)
	everySecond := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-inOneMinute.C:
			return
		case <-everySecond.C:
			// Yet another tick.
		}
	}
	// everySecond will never be GC'd!
	// END TICKER GC OMIT
}

func timerStale() {
	// BEGIN TIMER STOP OMIT
	t := time.NewTimer(1 * time.Minute)
	go func() {
		<-t.C
	}()
	if !t.Stop() { // HL
		<-t.C // HL
	} // HL
	t.Reset(1 * time.Second)
	// END TIMER STOP OMIT
}
