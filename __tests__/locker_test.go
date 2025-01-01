package __tests__

import (
	"sync"
	"testing"

	"github.com/udfordria/go-locker"
)

func TestLockerRead(t *testing.T) {
	// Create a Locker with an initial value
	lock := locker.NewLocker(42)

	// Read the value
	got := lock.Read()
	want := 42

	if got != want {
		t.Errorf("Read() = %v; want %v", got, want)
	}
}

func TestLockerSet(t *testing.T) {
	// Create a Locker with an initial value
	lock := locker.NewLocker(42)

	// Define the callback function that changes the value
	cb := func(val *int) *int {
		newVal := *val + 10
		return &newVal
	}

	// Call Set with the callback function
	lock.Set(cb)

	// Read the updated value
	got := lock.Read()
	want := 52

	if got != want {
		t.Errorf("Set() updated value to %v; want %v", got, want)
	}
}

func TestLockerSetNilCallback(t *testing.T) {
	// Create a Locker with an initial value
	lock := locker.NewLocker(42)

	// Define a callback that returns nil
	cb := func(val *int) *int {
		return nil
	}

	// Call Set with the nil callback
	lock.Set(cb)

	// Value should remain unchanged
	got := lock.Read()
	want := 42

	if got != want {
		t.Errorf("Set() with nil callback didn't leave value unchanged: got %v; want %v", got, want)
	}
}

func TestLockerConcurrentReadWrite(t *testing.T) {
	// Create a Locker with an initial value
	lock := locker.NewLocker(42)

	// Use a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Test concurrent reading and writing
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Perform a read
			lock.Read()

			// Perform a write (set the value)
			lock.Set(func(val *int) *int {
				newVal := *val + 1
				return &newVal
			})
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// After 100 iterations, the final value should be 142
	got := lock.Read()
	want := 142

	if got != want {
		t.Errorf("Expected value after 100 writes: %v, got %v", want, got)
	}
}
