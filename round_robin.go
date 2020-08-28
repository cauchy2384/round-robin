package roundrobin

import (
	"fmt"
	"math"
	"sync/atomic"
)

// RoundRobin provides strings from list via method Next() in given order and thread-safe way.
type RoundRobin struct {
	list    []string
	listLen uint64
	counter uint64
}

// New RoundRobin instance with given list.
// Returns ErrorInvalidConfig on error.
func New(list []string) (*RoundRobin, error) {
	if list == nil {
		return nil, fmt.Errorf("%w: list is empty", ErrorInvalidConfig)
	}

	rr := RoundRobin{
		list:    list,
		listLen: uint64(len(list)),
		counter: math.MaxUint64, // Next() should return element with idx 0 on first call
	}

	return &rr, nil
}

// Next string from list is returned to caller.
func (rr *RoundRobin) Next() string {
	counter := atomic.AddUint64(&rr.counter, 1)
	idx := counter % rr.listLen
	return rr.list[idx]
}
