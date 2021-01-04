package rollingwindow

import (
	"fmt"
	"time"
)

const duration = 100 * time.Millisecond

func ExampleRollingWindow_Add() {
	const size = 3
	r := NewRollingWindow(size, duration)
	listBuckets := func() []int64 {
		var buckets []int64
		r.Reduce(func(b *bucket) {
			buckets = append(buckets, b.value)
		})
		return buckets
	}
	r.Add(1)
	time.Sleep(duration)
	r.Add(2)
	time.Sleep(duration)
	r.Add(3)
	fmt.Println(listBuckets())
	time.Sleep(3 * duration)
	fmt.Println(listBuckets())
	// Output:
	// [1 2 3]
	// []
}
