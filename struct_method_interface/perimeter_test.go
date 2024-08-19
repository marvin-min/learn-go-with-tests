package struct_method_interface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	assetEqual(got, want, t)
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})
	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})

	t.Run("table test", func(t *testing.T) {
		areaTest := []struct {
			shape Shape
			want  float64
		}{
			{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
			{shape: Circle{Radius: 10}, want: 314.1592653589793},
			{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
		}
		for _, tt := range areaTest {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		}
	})
}

func assetEqual(got, want float64, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %.g want %.g", got, want)
	}
}
