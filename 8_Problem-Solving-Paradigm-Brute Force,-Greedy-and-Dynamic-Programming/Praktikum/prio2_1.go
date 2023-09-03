package main
import (
	"fmt"
	"math"
)
func Frog(jumps []int) int {
	n := len(jumps)
	if n < 2 {
		return 0
	}
	
	totalbiaya := make([]int, n)
	totalbiaya[0] = 0
	totalbiaya[1] = int(math.Abs(float64(jumps[1] - jumps[0])))
	
	for i := 2; i < n; i++ {
		totalbiaya[i] = min(totalbiaya[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1]))), totalbiaya[i-2]+int(math.Abs(float64(jumps[i]-jumps[i-2]))))
	}

	return totalbiaya[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
