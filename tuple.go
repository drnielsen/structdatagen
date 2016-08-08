package main

import "sync"

// Tuple stores multiple key/value pairs
type Tuple struct {
	keyValuePairs map[string]interface{}
	sync.RWMutex
}

// NewTuple creates a new Tuple and returns it.
func NewTuple() *Tuple {
	var t Tuple
	t.keyValuePairs = make(map[string]interface{})
	return &t
}

// Get the value for the specified column
func (t *Tuple) Get(columnName string) interface{} {
	t.RLock()
	value := t.keyValuePairs[columnName]
	t.RUnlock()
	return value
}

// Set the value in the specified column
func (t *Tuple) Set(columnName string, value interface{}) {
	t.Lock()
	t.keyValuePairs[columnName] = value
	t.Unlock()
}

// GetColNames will return a slice of all column
// names in a Tuple
func (t *Tuple) GetColNames() []string {
	t.RLock()
	keys := make([]string, len(t.keyValuePairs))

	i := 0
	for k := range t.keyValuePairs {
		keys[i] = k
		i++
	}
	t.Unlock()

	return keys
}
