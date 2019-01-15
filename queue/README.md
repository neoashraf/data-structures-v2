# Queue
the queue will be implemented with `slice` type.

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

