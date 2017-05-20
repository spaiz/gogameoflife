package tests

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spaiz/gogameoflife"
	"strings"
)

func TestMultilineInputDataParsing(t *testing.T) {
	Convey("Given valid multiline input string data, should return NxN matrix", t, func() {
		input := []string{
			"-----",
			"OOOOO",
			"-----",
			"--O--",
			"-----",
		}

		expected := gogameoflife.Desk{
			{false, false, false, false, false},
			{true, true, true, true, true},
			{false, false, false, false, false},
			{false, false, true, false, false},
			{false, false, false, false, false},
		}

		actual, err := gogameoflife.ParseInputData(strings.Join(input, "\n"), "O")
		So(err, ShouldBeNil)
		So(actual, ShouldResemble, expected)
	})

	Convey("Given non aligned multiline input string data, should return error", t, func() {
		input := []string{
			"-----",
			"-",
			"-----",
			"-----",
			"-----",
		}

		actual, err := gogameoflife.ParseInputData(strings.Join(input, "\n"), "O")
		So(err, ShouldNotBeNil)
		So(actual, ShouldBeNil)
	})

	Convey("Given empty input string data, should return error", t, func() {
		actual, err := gogameoflife.ParseInputData("", "O")
		So(err, ShouldNotBeNil)
		So(actual, ShouldBeNil)
	})
}