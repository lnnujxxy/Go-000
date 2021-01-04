package rollingwindow

import (
	"sync"
	"time"
)

type (
	RollingWindow struct {
		size     int
		win      *window
		interval time.Duration
		offset   int
		lastTime time.Time

		mu sync.RWMutex
	}

	Option func(rollingWindow *RollingWindow)
)

func NewRollingWindow(size int, interval time.Duration, opts ...Option) *RollingWindow{
	if size <= 0 {
		panic("size invalid")
	}
	w := &RollingWindow{
		size:     size,
		win:      newWindow(size),
		interval: interval,
		lastTime: time.Now(),
	}

	for _, opt := range opts {
		opt(w)
	}
	return w
}

func (rw *RollingWindow) Add(v int64) {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.updateOffset()
	rw.win.add(rw.offset, v)
}

func (rw *RollingWindow) Reduce(fn func(b *bucket)) {
	rw.mu.RLock()
	defer rw.mu.RUnlock()

	span := rw.span()
	diff := rw.size - span
	if diff > 0 {
		offset := (rw.offset + span + 1) % rw.size
		rw.win.reduce(offset, diff, fn)
	}
}

func (rw *RollingWindow) updateOffset() {
	span := rw.span()
	if span <= 0 {
		return
	}

	offset := rw.offset
	start := offset + 1
	steps := start + span
	remainder := 0

	if steps > rw.size {
		remainder = steps - rw.size
		steps = rw.size
	}

	for i:= start; i < steps; i++ {
		rw.win.reset(i)
	}
	for i := 0; i < remainder; i ++ {
		rw.win.reset(i)
	}

	rw.offset= (offset + span) % rw.size
	rw.lastTime = time.Now()
}

func (rw *RollingWindow) span() int {
	offset := int(time.Since(rw.lastTime)/rw.interval)
	if 0 <= offset && offset < rw.size {
		return offset
	}
	return rw.size
}

type bucket struct {
	value int64
	count int64
}

func (b *bucket) add(v int64) {
	b.value += v
	b.count++
}

func (b *bucket) reset() {
	b.value = 0
	b.count = 0
}

type window struct {
	buckets []*bucket
	size int
}

func newWindow(size int) *window {
	buckets := make([]*bucket, size)
	for i := 0; i < size; i++ {
		buckets[i] = new(bucket)
	}

	return &window{
		buckets: buckets,
		size:    size,
	}
}

func (w *window) add(offset int, v int64) {
	w.buckets[offset%w.size].add(v)
}

func (w *window) reset(offset int) {
	w.buckets[offset%w.size].reset()
}

func (w *window) reduce(start, count int, fn func(b *bucket)) {
	for i := 0; i < count; i++ {
		fn(w.buckets[(start+i)%w.size])
	}
}