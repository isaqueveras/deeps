package core

import (
	"context"
	"time"

	"github.com/isaqueveras/leaf"
)

// Queue ...
type Queue struct {
	q leaf.IQueue
}

// NewQueue ...
func NewQueue() *Queue {
	return &Queue{q: leaf.New(context.Background(), 100, 100, time.Second/100)}
}

// PublishFunc ...
type PublishFunc func(context.Context) (interface{}, error)

// Publish ...
func (q *Queue) Publish(fn PublishFunc) {
	q.q.Publish(func(ctx context.Context) (interface{}, error) {
		return fn(ctx)
	})
}

func (q *Queue) Stop(ctx context.Context) {
	leaf.Stop(ctx)
}

type page leaf.IPage

func (q *Queue) GetPage(ctx context.Context) page {
	return leaf.GetPage(ctx)
}

// ConsumerFunc ...
type ConsumerFunc func(ctx context.Context, data interface{}) error

// Consume defines the method to consume data from the queue
func (q *Queue) Consume(fn ConsumerFunc) {
	q.q.Consume(func(ctx context.Context) error {
		return fn(ctx, leaf.GetData(ctx))
	})
}

// Wait defines the method to wait for
// publishers and consumers to execute
func (q *Queue) Wait() {
	q.q.Wait()
}
