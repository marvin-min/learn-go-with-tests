package generics

type MyIntStack struct {
	values []int
}

func (mis *MyIntStack) Push(value int) {
	mis.values = append(mis.values, value)
}

func (mis *MyIntStack) Pop() (int, bool) {
	if mis.IsEmpty() {
		return 0, false
	}
	index := len(mis.values) - 1
	el := mis.values[index]
	mis.values = mis.values[:index]
	return el, true
}

func (mis *MyIntStack) IsEmpty() bool {
	return len(mis.values) == 0
}
