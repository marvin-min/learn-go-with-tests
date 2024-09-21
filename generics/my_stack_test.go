package generics_test

import (
	"testing"

	"github.com/marvin-min/learn-go-with-tests/generics"
)

func TestMyIntStack(t *testing.T) {
	t.Run("int stack operations", func(t *testing.T) {
		mis := new(generics.Stack[int])
		AssertTrue(t, mis.IsEmpty())

		mis.Push(1)
		mis.Push(2)
		mis.Push(3)
		mis.Push(123)
		AssertFalse(t, mis.IsEmpty())
		value, _ := mis.Pop()
		AssertEqual(t, value, 123)
		value1, _ := mis.Pop()
		AssertEqual(t, value1, 3)

		mis.Pop()
		mis.Pop()
		AssertTrue(t, mis.IsEmpty())
	})

	t.Run("string stack operations", func(t *testing.T) {
		mis := new(generics.Stack[string])
		AssertTrue(t, mis.IsEmpty())

		mis.Push("1")
		mis.Push("2")
		mis.Push("3")
		mis.Push("123")
		AssertFalse(t, mis.IsEmpty())
		value, _ := mis.Pop()
		AssertEqual(t, value, "123")
		value1, _ := mis.Pop()
		AssertEqual(t, value1, "3")

		mis.Pop()
		mis.Pop()
		AssertTrue(t, mis.IsEmpty())
	})
}

func AssertTrue(t testing.TB, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t testing.TB, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
