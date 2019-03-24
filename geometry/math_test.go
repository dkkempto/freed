package geometry

import "testing"

func TestAdd(t *testing.T) {
	res := Add([3]float64{0, 0, 0}, [3]float64{1, 1, 1}, [3]float64{9, 8, 7})

	if res[0] != 10 {
		t.Errorf("X was incorrect - got: %f, want: %f", res[0], 10.0)
	}
	if res[1] != 9 {
		t.Errorf("Y was incorrect - got: %f, want: %f", res[1], 9.0)
	}
	if res[2] != 8 {
		t.Errorf("Z was incorrect - got: %f, want: %f", res[2], 8.0)
	}
}

func TestDot(t *testing.T) {
	res := Dot([3]float64{1, 2, 3}, [3]float64{4, 5, 6})

	if res != 32 {
		t.Errorf("Error calculating dot product - got: %f, want: %f", res, 32.0)
	}
}

func TestCross(t *testing.T) {
	res := Cross([3]float64{1, 2, 3}, [3]float64{4, 5, 6})

	if res[0] != -3 {
		t.Errorf("X was incorrect - got: %f, want: %f", res[0], -3.0)
	}

	if res[1] != 6 {
		t.Errorf("Y was incorrect - got: %f, want: %f", res[1], 6.0)
	}

	if res[2] != -3 {
		t.Errorf("Z was incorrect - got: %f, want: %f", res[2], -3.0)
	}
}
