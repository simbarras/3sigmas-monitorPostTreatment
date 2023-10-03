package equation

import "testing"

func TestAddition_Compute(t *testing.T) {
	add := Addition{}

	if add.Compute(1, 2) != 3 {
		t.Error("Addition failed: 1 + 2 != 3")
	}

	if add.Compute(0, 1) != 1 {
		t.Error("Addition failed: 0 + 1 != 1")
	}

	if add.Compute(-1, 1) != 0 {
		t.Error("Addition failed: -1 + 1 != 0")
	}
}

func FuzzAddition_Compute(f *testing.F) {
	f.Add(0, 1)
	f.Add(1, 2)
	f.Add(-1, 1)
	f.Fuzz(func(t *testing.T, a int, b int) {
		add := Addition{}
		if add.Compute(a, b) != a+b {
			t.Error("Addition failed: a + b != a + b")
		}
	})
}
