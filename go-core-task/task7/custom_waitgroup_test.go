package main

import (
	"sync/atomic"
	"testing"
)

func TestCustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup()
	counter := int64(0)
	totalTasks := int64(5)

	for i := int64(0); i <= totalTasks; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()

	if counter != totalTasks {
		t.Errorf("Expected counter to be %d, got %d", totalTasks, counter)
	}
}