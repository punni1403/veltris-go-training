package main

import "testing"

type MockAdder struct{}

func (m *MockAdder) Add(a, b int) int {
	return 5
}

func TestAddMock(t *testing.T) {

	// adder = &MockAdder{}
	// adder.Add()

	res := Add(10, 2)

	if res != 20 {
		t.Errorf("Expected: %d, got = %d", 20, res)
	}

}

func Test_Add(t *testing.T) {

	// res := Add(2, 5)
	// expecting := 5

	// if res != expecting {
	// 	t.Errorf("Expected: %d, got = %d", expecting, res)
	// }

	// t.Run("Add positive no", func(t *testing.T) {

	// })

	tests := []struct {
		a, b, expecting int
	}{
		{2, 3, 5},
		{-2, -3, -5},
		{0, 0, 0},
		{1, 1, 2},
	}

	for _, tt := range tests {
		res := Add(tt.a, tt.b)
		if res != tt.expecting {
			t.Errorf("Expected: %d, got = %d", tt.expecting, res)
		}
	}

}

// Benchmark-ing

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}
