// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Package queue creates a StringQueue data structure for the string type
package main

import (
	"fmt"
	"sync"
)

// StringQueue the queue of Strings
type StringQueue struct {
	items []string
	lock  sync.RWMutex
}

// New creates a new StringQueue
func (s *StringQueue) New() *StringQueue {
	s.items = []string{}
	return s
}

// Enqueue adds an string to the end of the queue
func (s *StringQueue) Enqueue(t string) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Dequeue removes an string from the start of the queue
func (s *StringQueue) Dequeue() *string {
	s.lock.Lock()
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	s.lock.Unlock()
	return &item
}

// Front returns the item next in the queue, without removing it
func (s *StringQueue) Front() *string {
	s.lock.RLock()
	item := s.items[0]
	s.lock.RUnlock()
	return &item
}

// IsEmpty returns true if the queue is empty
func (s *StringQueue) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of Strings in the queue
func (s *StringQueue) Size() int {
	return len(s.items)
}

func main() {
	s := &StringQueue{}

	s.Enqueue("A")
	s.Enqueue("B")
	s.Enqueue("C")

	for !s.IsEmpty() {
		fmt.Println("Front", "=>", *s.Front())
		fmt.Println(*s.Dequeue())
		fmt.Println("Front", "=>", *s.Front())
		fmt.Println(*s.Dequeue())
		fmt.Println("Front", "=>", *s.Front())
		fmt.Println(*s.Dequeue())
	}
}
