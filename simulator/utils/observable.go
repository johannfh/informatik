package utils

import "log/slog"

type Observable[T any] struct {
	value T

	listeners []ListenerFunc[T]
}

type ListenerFunc[T any] func(val T)

func (o *Observable[T]) Get() T {
	return o.value
}

func (o *Observable[T]) Set(val T) {
	for _, fn := range o.listeners {
		slog.Info("called event listener")
		fn(val)
	}
	o.value = val
}

func (o *Observable[T]) OnChange(fn ListenerFunc[T]) {
	o.listeners = append(o.listeners, fn)
}
