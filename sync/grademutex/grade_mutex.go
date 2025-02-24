package grademutex

import (
	"sync"
)

type GradeMutex struct {
	read  sync.RWMutex
	write sync.Mutex
}

func (m *GradeMutex) Lock() {
	m.write.Lock()
	m.read.Lock()
}

func (m *GradeMutex) Unlock() {
	m.read.Unlock()
	m.write.Unlock()
}

func (m *GradeMutex) RLock() {
	m.write.Lock()
	m.read.RLock()
	m.write.Unlock()
}

func (m *GradeMutex) RUnlock() {
	m.read.RUnlock()
}

func (m *GradeMutex) Upgrade() {
	m.write.Lock()
	m.read.RUnlock()
	m.read.Lock()
}

func (m *GradeMutex) Downgrade() {
	m.read.Unlock()
	m.read.RLock()
	m.write.Unlock()
}
