# locker

Provide a secured way to use mutex in go.

## Installation

```cmd
go get -u github.com/udfordria/go-locker
```

## Usage

```go
import (
	"github.com/udfordria/go-locker"
)

func main() {
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
```
