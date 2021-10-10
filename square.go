package parker

// NumberSquare represents numbers in the (3x3) square using the following
// convention for indeces:
// 0 1 2
// 3 4 5
// 6 7 8
type NumberSquare struct {
	values [9]int
}

func (ps *NumberSquare) Set(idx, value int) {
	ps.values[idx] = value
}

func (ps *NumberSquare) IsParker() bool {
	return true
}

func (ps *NumberSquare) SquareValues() {
	for idx, value := range ps.values {
		ps.values[idx] = value * value
	}
}

func addOrIncrement(occurences map[int]int, value int) {
	if occurence, ok := occurences[value]; ok {
		occurences[value] = occurence + 1
	} else {
		occurences[value] = 1
	}
}

func biggestValue(occurences map[int]int) int {
	biggest := 0
	for _, value := range occurences {
		if value > biggest {
			biggest = value
		}
	}
	return biggest
}

func (ps *NumberSquare) SameValues() int {
	occurences := make(map[int]int)
	addOrIncrement(occurences, ps.rowSum(0))
	addOrIncrement(occurences, ps.rowSum(1))
	addOrIncrement(occurences, ps.rowSum(2))
	addOrIncrement(occurences, ps.colSum(0))
	addOrIncrement(occurences, ps.colSum(1))
	addOrIncrement(occurences, ps.colSum(2))
	addOrIncrement(occurences, ps.diagSum1())
	addOrIncrement(occurences, ps.diagSum2())
	return biggestValue(occurences)
}

func (ps *NumberSquare) rowSum(row int) int {
	return ps.values[3*row+0] + ps.values[3*row+1] + ps.values[3*row+2]
}

func (ps *NumberSquare) colSum(col int) int {
	return ps.values[1*col+0] + ps.values[1*col+3] + ps.values[1*col+6]
}

func (ps *NumberSquare) diagSum1() int {
	return ps.values[0] + ps.values[4] + ps.values[8]
}

func (ps *NumberSquare) diagSum2() int {
	return ps.values[2] + ps.values[4] + ps.values[6]
}
