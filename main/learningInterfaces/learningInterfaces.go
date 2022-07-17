package learningInterfaces

import "fmt"

type figure2D interface {
	area() float64
}

func calculate(f figure2D) {
	fmt.Printf("Area of: %v \n", f.area())
}

type square struct {
	base float64
}

func (s square) area() float64 {
	return s.base * s.base
}

type rectangle struct {
	base   float64
	height float64
}

func (r rectangle) area() float64 {
	return r.height * r.base
}

func LearningInterfaces() {
	mySquare := square{8}
	myRectangle := rectangle{
		base:   4,
		height: 6,
	}

	calculate(mySquare)
	calculate(myRectangle)
}
