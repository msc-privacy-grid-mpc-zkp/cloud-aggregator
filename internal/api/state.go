package api

import (
	"fmt"
	"sync"
)

type AggregationSession struct {
	TotalSum int64
	Count    int
	Meters   map[string]bool
}

type MemoryStore struct {
	mu             sync.Mutex
	Sessions       map[int64]*AggregationSession
	ExpectedMeters int
}

func NewMemoryStore(expected int) *MemoryStore {
	return &MemoryStore{
		Sessions:       make(map[int64]*AggregationSession),
		ExpectedMeters: expected,
	}
}

func (store *MemoryStore) AddShare(timestamp int64, meterID string, share int64) (bool, float64) {
	store.mu.Lock()
	defer store.mu.Unlock()

	session, exists := store.Sessions[timestamp]
	if !exists {
		session = &AggregationSession{
			TotalSum: 0,
			Count:    0,
			Meters:   make(map[string]bool),
		}
		store.Sessions[timestamp] = session
	}

	if session.Meters[meterID] {
		return false, 0
	}

	session.Meters[meterID] = true
	session.TotalSum += share
	session.Count++

	fmt.Printf("[AGGREGATOR] Progress for time %d: %d/%d meters\n", timestamp, session.Count, store.ExpectedMeters)

	if session.Count == store.ExpectedMeters {
		average := float64(session.TotalSum) / float64(store.ExpectedMeters)
		delete(store.Sessions, timestamp)

		return true, average
	}

	return false, 0
}
