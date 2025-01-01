package locker

import (
	"sync"
)

type Locker[A any] struct {
	m     *sync.RWMutex
	value A
}

func NewLocker[A any](initialValue A) Locker[A] {
	return Locker[A]{
		m:     &sync.RWMutex{},
		value: initialValue,
	}
}

func (rw *Locker[A]) Read() A {
	rw.m.RLock()
	defer rw.m.RUnlock()
	return rw.value
}

func (rw *Locker[A]) Set(cb func(*A) *A) {
	rw.m.RLock()
	res := cb(&rw.value)
	rw.m.RUnlock()
	if res != nil {
		rw.m.Lock()
		rw.value = *res
		rw.m.Unlock()
	}
}
