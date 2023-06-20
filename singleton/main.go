package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	data string
}

var instance *singleton
var mutex sync.Mutex

// getInstance returns the single instance of the singleton
func getInstance() *singleton {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			instance = &singleton{data: "Singleton Instance"}
		}
	}
	return instance
}

func main() {
	s1 := getInstance()
	s2 := getInstance()

	// Both instances are the same
	fmt.Println(s1.data) // Output: Singleton Instance
	fmt.Println(s2.data) // Output: Singleton Instance

	// Modifying data in one instance affects the other instance
	s1.data = "Modified Singleton Instance"
	fmt.Println(s1.data) // Output: Modified Singleton Instance
	fmt.Println(s2.data) // Output: Modified Singleton Instance
}
