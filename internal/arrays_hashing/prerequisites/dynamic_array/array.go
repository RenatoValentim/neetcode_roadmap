package dynamicarray

type ArrayInput interface {
	int | int64 | float64 | string
}

type array[T ArrayInput] struct {
	items  []T
	length int
}

func New[T ArrayInput]() *array[T] {
	return &array[T]{length: 10}
}

func (a *array[T]) Length() int {
	return a.length
}

func (a *array[T]) Add(item T) {
	a.items = append(a.items, item)
}

func (a *array[T]) Items() []T {
	return a.items
}
