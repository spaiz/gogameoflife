package gogameoflife

// Represents area where cells are live and die. Used as NxN matrix
type Desk [][]bool

// Find number of neighbours of specific cell
// Neighbour defined as live (true) cell around specific cell
func (r Desk) CountNeighbours(row int, column int) int {
	neighboursNum := 0

	for k := row -1; k <= row +1; k++ {
		for m := column - 1; m <= column +1; m++ {
			if k < 0 || k > len(r) - 1 || m < 0 || m > len(r) - 1 || (k == row && m == column)  {
				continue
			}

			if r[k][m] {
				neighboursNum++
			}
		}
	}

	return neighboursNum
}

// Decides if specific cell must become dead, born or stay unchanged
// in the next generation
func (r Desk) GetNewCellState(row int, column int) bool {
	neighbours := r.CountNeighbours(row, column)

	if r[row][column] {
		if neighbours < 2 || neighbours > 3 {
			return false
		}
		return r[row][column]
	} else if neighbours == 3 {
		return true
	}

	return false
}