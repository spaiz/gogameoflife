package tests

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spaiz/gogameoflife"
	"bytes"
)

func TestGame(t *testing.T) {
	Convey("Should write rendered desk's data to an writer", t, func() {
		var b bytes.Buffer

		desk := gogameoflife.Desk{
			{true, true, true, true},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}

		game := &gogameoflife.Game{CurrentGenerationDesk: desk}
		game.Render(&b)

		So(b.String(), ShouldEqual, "⠉⠉")
	})
}