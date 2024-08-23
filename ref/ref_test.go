package ref

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	// t.Run("basic walk", func(t *testing.T) {
	// 	expected := "Chris"
	// 	var got []string
	// 	x := struct {
	// 		Name string
	// 	}{expected}
	// 	walk(x, func(input string) {
	// 		got = append(got, input)
	// 	})

	// 	if len(got) != 1 {
	// 		t.Errorf("wrong number of arguments, got %d want 1", len(got))
	// 	}
	// 	if got[0] != expected {
	// 		t.Errorf("got %q, want %q", got[0], expected)
	// 	}
	// })

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"struct with one string field",
			struct {
				Name string
			}{"Chris"}, []string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				Age  int
			}{"Chris", 20},
			[]string{"Chris"},
		},
		{
			"struct with struct fields",
			Person{"Chris", Profile{20, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
