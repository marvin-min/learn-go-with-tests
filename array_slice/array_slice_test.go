package array_slice

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestArraySum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9, 9})
	want := []int{3, 18}
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 4, 5, 6, 7, 8, 9, 10}, []int{0, 9, 9, 8, 8, 3, 4, 5, 6, 6})
	}
}

func TestSumAllTails(t *testing.T) {
	assetEquals := func(got []int, want []int, t *testing.T) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assetEquals(got, want, t)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
