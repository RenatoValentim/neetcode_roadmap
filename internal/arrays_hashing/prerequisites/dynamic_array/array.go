package dynamicarray

type array struct {
	length int
}

func New() *array {
	return &array{length: 10}
}

func (a *array) Length() int {
	return a.length
}
