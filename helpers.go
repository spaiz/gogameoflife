package gogameoflife

import (
	"strings"
	"errors"
)

var ErrorStringDataMustBeAligned = errors.New("Please, provide initial data as NxN struture where N > 0. Number of chars in each line must be the same as number of lines.")
var ErrorEmptyInput = errors.New("Empty input data received by parser")

// Creates matrix NxN and fills it with booleans related to provided multiline string
// By convention, data might be aligned.
func ParseInputData(str string, liveSymbol string) (Desk, error) {
	var matrix Desk

	strLines := strings.Split(strings.Trim(str, "\n"), "\n")

	if str == "" {
		return nil, ErrorEmptyInput
	}

	if !IsValidInput(strLines) {
		return nil, ErrorStringDataMustBeAligned
	}

	matrixSize := len(strLines)
	matrix = CreateDefaultDeadDesk(matrixSize)

	for i := range strLines {
		for y := 0; y < matrixSize; y++ {
			if string(strLines[i][y]) == liveSymbol {
				matrix[i][y] = true
			}
		}
	}

	return matrix, nil
}

// We allow only good source data
// Number of chars in each line must be equal to number of lines
func IsValidInput(strLines []string) bool {
	linesNum := len(strLines)

	for i := range strLines {
		if len(strLines[i]) != linesNum {
			return false
		}
	}

	return true
}

// Creates default square matrix filled with all false booleans
func CreateDefaultDeadDesk(size int) Desk {
	matrix := make(Desk, size)

	for i := range matrix {
		matrix[i] = make([]bool, size)
	}

	return matrix
}