package singleton

import (
	"sync"
	"testing"
)

func TestGetInstance_SingleInstanceCreated(t *testing.T) {
	// Reset the singleton instance for testing purposes
	singleInstance = nil

	// First call: should create a new instance
	instance1 := GetInstance()
	if instance1 == nil {
		t.Error("Expected an instance, got nil")
	}

	// Second call: should return the existing instance
	instance2 := GetInstance()
	if instance1 != instance2 {
		t.Error("Expected the same instance, got a different one")
	}
}

func TestGetInstance_ConcurrentAccess(t *testing.T) {
	// Reset the singleton instance for testing purposes
	singleInstance = nil

	var wg sync.WaitGroup
	instanceSet := make(map[*single]bool)
	mu := sync.Mutex{}

	// Simulate concurrent access to GetInstance
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			instance := GetInstance()

			mu.Lock()
			instanceSet[instance] = true
			mu.Unlock()
		}()
	}

	wg.Wait()

	if len(instanceSet) != 1 {
		t.Errorf("Expected only one instance, but got %d", len(instanceSet))
	}
}
