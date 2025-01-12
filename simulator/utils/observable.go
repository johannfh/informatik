package utils

type Observable[T any] struct {
	value T

	listeners []ListenerFunc[T]
}

func NewObservable[T any](val T) Observable[T] {
	return Observable[T]{
		value:     val,
		listeners: []ListenerFunc[T]{},
	}
}

type ListenerFunc[T any] func(val T)

func (o *Observable[T]) Get() T {
	return o.value
}

func (o *Observable[T]) Set(val T) {
	o.value = val
	o.Trigger()
}

func (o *Observable[T]) Trigger() {
	for _, fn := range o.listeners {
		fn(o.value)
	}
}

func (o *Observable[T]) OnChange(fn ListenerFunc[T]) {
	o.listeners = append(o.listeners, fn)
}
