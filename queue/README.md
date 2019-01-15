# Queue
the queue will be implemented with `slice` type.

## Installation
```bash
go get github.com/cheekybits/genny
cd <this directory>
cat queue.go | genny gen "Item=string"
```

## Methods
- New()
- Enqueue()
- Dequeue()
- Front()
- IsEmpty()
- Size()

### Base Type
`ItemQueue` Generic type, concurrency safe, that can generate queue of any type by using [genny](https://github.com/cheekybits/genny).

``` go
package queue
import (
    "sync"

    "github.com/cheekybits/genny/generic"
)

// Item the type of the queue
type Item generic.Type

// ItemQueue the queue of Items
type ItemQueue struct {
    items []Item
    lock  sync.RWMutex
}

```

### New() 
* constructor
* initializes the slice

``` go
// New creates a new ItemQueue
func (s *ItemQueue) New() *ItemQueue {
    s.items = []Item{}
    return s
}
```

### Enqueue()
``` go
// Enqueue adds an Item to the end of the queue
func (s *ItemQueue) Enqueue(t Item) {
    s.lock.Lock()
    s.items = append(s.items, t)
    s.lock.Unlock()
}
```

### Dequeue()
``` go
// Dequeue removes an Item from the start of the queue
func (s *ItemQueue) Dequeue *Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return &item
}
```

### Front()
``` go
// Front returns the item next in the queue w/o removing it
func (s *ItemQueue) Front() *Item {
	s.lock.RLock()
	defer s.lock.RUnlock()
	item := s.items[0]
	return &item
}
```

### IsEmpty()

``` go
// IsEmpty returns true if the queue is empty
func (s *ItemQueue) IsEmpty() bool {
    return len(s.items) == 0
}

```

### Size()

``` go
// Size returns the number of Items in the queue
func (s *ItemQueue) Size() int {
    return len(s.items)
}
```

