package generics

// type MyIntStack = Stack
// type MyStringStack = Stack
type Stack[T any] struct {
	values []T
}

func (mis *Stack[T]) Push(value T) {
	mis.values = append(mis.values, value)
}

func (mis *Stack[T]) Pop() (T, bool) {
	if mis.IsEmpty() {
		var zero T
		return zero, false
	}
	index := len(mis.values) - 1
	el := mis.values[index]
	mis.values = mis.values[:index]
	return el, true
}

func (mis *Stack[T]) IsEmpty() bool {
	return len(mis.values) == 0
}
