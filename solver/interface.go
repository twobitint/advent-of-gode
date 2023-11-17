package solver

type Solver interface {
	PartOne() any
	PartTwo() any

	SetInput(string)
	GetInput() string
}

type Solves struct {
	input string
}

func (s Solves) PartOne() any {
	return "Part 1 not implemented"
}

func (s Solves) PartTwo() any {
	return "Part 2 not implemented"
}

func (s *Solves) SetInput(input string) {
	s.input = input
}

func (s Solves) GetInput() string {
	return s.input
}
