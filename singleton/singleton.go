package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{} // Mutex to ensure thread-safe access to the singleton instance.

type single struct{} // Struct representing the singleton instance.

var singleInstance *single // Variable to hold the singleton instance.

// GetInstance provides a globally accessible method to get the singleton instance.
// Ensures that only one instance of 'single' is created, even in a concurrent environment.
func GetInstance() *single {
	if singleInstance == nil { // Check if the singleton instance is not yet created.
		lock.Lock() // Lock to ensure only one goroutine can create the instance.
		defer lock.Unlock()
		fmt.Println("new instance being created")
		singleInstance = &single{} // Create the singleton instance.
	} else {
		fmt.Println("existing instance being returned")
	}
	return singleInstance // Return the singleton instance.
}
