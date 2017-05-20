package tests

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spaiz/gogameoflife"
)

func TestMultilineFileInputDataProvider(t *testing.T) {
	Convey("Given an existing file path, file's content should be returned", t, func() {
		filePath := "./fixtures/valid_input.txt"

		var provider gogameoflife.DataSourceProvider = gogameoflife.NewFileSourceDataProvider(filePath)

		content, err := provider.GetData()

		So(err, ShouldBeNil)
		So(content, ShouldNotBeEmpty)
		So(provider.GetLiveSymbol(), ShouldNotBeEmpty)
	})

	Convey("Given an non existing file path, an error should be returned", t, func() {
		filePath := "./fixtures/nonexisting_path_file.txt"

		var provider gogameoflife.DataSourceProvider = gogameoflife.NewFileSourceDataProvider(filePath)

		content, err := provider.GetData()
		So(err, ShouldNotBeNil)
		So(content, ShouldBeEmpty)
	})
}

func TestMultilineRandomInputDataProvider(t *testing.T) {
	Convey("With valid size provided, should return random multiline string input data", t, func() {
		size := 10
		density := 3
		provider := gogameoflife.NewRandomSourceDataProvider(size, density)
		str, err := provider.GetData()
		So(err, ShouldBeNil)
		So(str, ShouldNotBeEmpty)
	})

	Convey("With invalid size provided, should return error", t, func() {
		size := -1
		density := 3
		provider := gogameoflife.NewRandomSourceDataProvider(size, density)
		str, err := provider.GetData()
		So(err, ShouldNotBeNil)
		So(str, ShouldBeEmpty)
	})
}