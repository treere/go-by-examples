package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 5}
	got := Perimeter(rectangle)
	want := 30.0

	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{10, 5}, want: 50.0},
		{name: "circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v: got '%g' want '%g'", tt.name, got, tt.want)
		}
	}
}
