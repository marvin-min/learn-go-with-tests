package generics

type MyIntStack = Stack
type MyStringStack = Stack
type Stack struct {
	values []interface{}
}

func (mis *MyIntStack) Push(value interface{}) {
	mis.values = append(mis.values, value)
}

func (mis *MyIntStack) Pop() (interface{}, bool) {
	if mis.IsEmpty() {
		var zero interface{}
		return zero, false
	}
	index := len(mis.values) - 1
	el := mis.values[index]
	mis.values = mis.values[:index]
	return el, true
}

func (mis *MyIntStack) IsEmpty() bool {
	return len(mis.values) == 0
}
