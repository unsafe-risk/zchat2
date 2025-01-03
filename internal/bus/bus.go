package bus

import "sync"

type optype uint8

const (
	optypeUnknown optype = iota
	optypeSubscribe
	optypeUnsubscribe
	optypePublish
)

type Topic[T any] struct {
	topic string
	subs  []*Subscription[T]
	mu    sync.Mutex
}

type Subscription[T any] struct {
	b  *Bus[T]
	t  *Topic[T]
	ch chan T
}

type Bus[T any] struct {
	subscribers map[string]*Topic[T]
	stopCh      chan struct{}
	opPool      sync.Pool
}

type operation[T any] struct {
	optype optype
	data   T
}
