package tests

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spaiz/gogameoflife"
)

func TestGameGeneration(t *testing.T) {
	Convey("Given an desk and position (i row, j column), next cell's state should be returned", t, func() {
		desk := gogameoflife.Desk{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, true, true, true, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}

		actual := desk.GetNewCellState(1, 2)
		So(actual, ShouldBeTrue)

		actual = desk.GetNewCellState(2, 1)
		So(actual, ShouldBeFalse)
	})

	Convey("Given an desk and position (i row, j column) the neighbours number should be returned", t, func() {
		desk := gogameoflife.Desk{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, true, true, true, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}

		neighbours := desk.CountNeighbours(1, 2)
		So(neighbours, ShouldEqual, 3)

		neighbours = desk.CountNeighbours(2, 1)
		So(neighbours, ShouldEqual, 1)
	})

	Convey("Given an desk, next population generation should be created", t, func() {
		current := gogameoflife.Desk{
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
		}

		expected := gogameoflife.Desk{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, true, true, true, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}

		game := &gogameoflife.Game{CurrentGenerationDesk: current}
		actual := game.NextGeneration()
		So(actual, ShouldResemble, expected)
	})

	Convey("Given desk size N, matrix of size NxN should be returned", t, func() {
		expected := gogameoflife.Desk{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}

		deskSize := 3
		desk := gogameoflife.CreateDefaultDeadDesk(deskSize)
		So(desk, ShouldResemble, expected)
	})
}