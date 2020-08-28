package roundrobin

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		list        []string
		expectedErr error
	}{
		{nil, ErrorInvalidConfig},
		{[]string{"192.168.1.1"}, nil},
	}

	for _, tt := range testCases {
		_, err := New(tt.list)
		if !errors.Is(err, tt.expectedErr) {
			t.Fatal("unexpected error", err, "should be", tt.expectedErr)
		}
	}
}

func TestRoundRobin_Next(t *testing.T) {
	list := []string{"Frodo", "Samwise", "Meriadoc", "Peregrin"}

	rr, err := New(list)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	for c := 0; c < 2; c++ { // cycle list twice
		for i, name := range list {
			nextValue := rr.Next()
			if nextValue != name {
				t.Fatal("expected", i, "to be", name, "not", nextValue)
			}
			_, _ = fmt.Fprintln(ioutil.Discard, nextValue)
		}
	}
}

func TestRoundRobin_NextRace(t *testing.T) {
	list := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4"}

	rr, err := New(list)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	// expect it to be enough for context switches and data racing
	const (
		consumersNum = 10
		requestsNum  = 100
	)
	wg := sync.WaitGroup{}
	wg.Add(consumersNum)

	for i := 0; i < consumersNum; i++ {
		go func(consumerNum int) {
			defer wg.Done()
			for j := 0; j < requestsNum; j++ {
				host := rr.Next()
				_, _ = fmt.Fprintln(ioutil.Discard, "consumer", consumerNum, "received:", host)
			}
		}(i)
	}

	wg.Wait()
}
