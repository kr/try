// Package try provides a non-blocking mutex.
package try

import (
	"sync"
)

// Mutex is a non-blocking mutual exclusion lock. Attempts
// to lock don't wait for the lock to be released. The zero
// value for Mutex is an unlocked mutex.
type Mutex struct {
	m      sync.Mutex
	locked bool
}

// TryLock attempts to acquire the lock without waiting
// for other holders to release it. It returns whether
// the attempt was successful.
func (t *Mutex) TryLock() bool {
	t.m.Lock()
	defer t.m.Unlock()
	if t.locked {
		return false
	}
	t.locked = true
	return true
}

// Unlock unlocks the mutex. It panics if the mutex is
// not locked when Unlock is called.
func (t *Mutex) Unlock() {
	t.m.Lock()
	defer t.m.Unlock()
	if !t.locked {
		panic("double call to unlock")
	}
	t.locked = false
}
