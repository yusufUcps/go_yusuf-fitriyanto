package main
import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	totalnilai:= 0
	for _, nilai := range s.score {
		totalnilai += nilai
	}
	return float64(totalnilai) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for a := 1; a < len(s.score); a++ {
		if s.score[a] < min {
			min = s.score[a]
			name = s.name[a]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for a := 1; a < len(s.score); a++ {
		if s.score[a] > max {
			max = s.score[a]
			name = s.name[a]
		}
	}
	return max, name
}
func main() {
	var a = Student{}
	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Input Student ", i+1, "'s Name: ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input ", name, "'s Score: ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}
	fmt.Println("\nAverage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Student is:", nameMax, "(", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Student is:", nameMin, "(", scoreMin, ")")
}