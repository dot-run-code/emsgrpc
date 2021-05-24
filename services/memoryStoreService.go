package services

import "sync"

type MemoryStore struct {
	store map[string]Event
	mu    sync.RWMutex
}

var instance *MemoryStore
var once sync.Once

func GetOrCreateMemoryStore() *MemoryStore {
	once.Do(func() {
		s := make(map[string]Event)
		instance = &MemoryStore{store: s}
	})
	return instance
}

func (m *MemoryStore) IsTopicRegistered(topic string) bool {
	m.mu.RLock()
	_, ok := m.store[topic]
	m.mu.RUnlock()
	if ok {
		return true
	}
	return false
}
func (m *MemoryStore) AddRemoveTopic(event Event) {
	m.mu.Lock()
	_, ok := m.store[event.Topic]
	if ok {
		delete(m.store, event.Topic)
	}

	switch event.Operation {
	case "Add":
		m.store[event.Topic] = event
	}
	m.mu.Unlock()
}
