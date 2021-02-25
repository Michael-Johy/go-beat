package mysum

import "testing"

func TestSum(t *testing.T) {

	cases := []struct {
		a, b int32
		c    int64
	}{
		{1, 2, 3},
		{2, 3, 2},
		{0, 0, 0},
	}

	for _, c := range cases {
		value := Sum(c.a, c.b)
		if value != c.c {
			t.Errorf("Sum(%d, %d) == %d, want = %d", c.a, c.b, value, c.c)
		}
	}

}
