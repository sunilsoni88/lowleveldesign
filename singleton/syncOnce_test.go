package singleton

import (
	"sync"
	"testing"
)

func TestGetSingleInstance_ResourceInitialization(t *testing.T) {
	// Reset the singleton instance for testing
	singleOnceInstance = nil
	once = sync.Once{}

	// First call: should initialize the resource
	instance1 := GetSingleInstance()
	if instance1 == nil {
		t.Error("Expected an instance, got nil")
	}
	if !instance1.resoureInitialized {
		t.Error("Expected resource to be initialized, but it was not")
	}

	// Second call: should return the already initialized instance
	instance2 := GetSingleInstance()
	if instance1 != instance2 {
		t.Error("Expected the same instance, got a different one")
	}
}

func TestGetSingleInstance_ConcurrentAccess(t *testing.T) {
	// Reset the singleton instance for testing
	singleOnceInstance = nil
	once = sync.Once{}

	var wg sync.WaitGroup
	instanceSet := make(map[*singleOnce]bool)
	mu := sync.Mutex{}

	// Simulate concurrent access to GetSingleInstance
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			instance := GetSingleInstance()

			mu.Lock()
			instanceSet[instance] = true
			mu.Unlock()
		}()
	}

	wg.Wait()

	// Only one instance should be created
	if len(instanceSet) != 1 {
		t.Errorf("Expected only one instance, but got %d", len(instanceSet))
	}
}
