package numbers

import "testing"

func BenchmarkNumberHard(b *testing.B) {
	// 952,25,50,75,100,3,6
	for n := 0; n < b.N; n++ {
		Solve([]int{25, 50, 75, 100, 3, 6}, 952)
	}
}
