package mutex

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	defer c.mu.Unlock()
}

func (c *Counter) Decrement() {
	c.mu.Lock()
	c.value--
	defer c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	// 읽기 시에도 락이 필요할 수 있다
	// 따라서 defer로 unlock..
	// 이 패턴 많이쓰일거같음..
	defer c.mu.Unlock()
	return c.value
}
