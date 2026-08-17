package counter

import (
	"sync"
)

type Counter struct{
	mu sync.Mutex
	hits int
}

func (c *Counter) Add(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.hits++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.hits
}
