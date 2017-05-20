// Use this provider when you want the game to load initial cells data from file

package gogameoflife

import (
	"io/ioutil"
	"math/rand"
	"errors"
	"time"
)

var ErrInvalidDeskSize = errors.New("Please provide positive non zero size for random data provider")

// Use it when you want to use custom source for Game's initial data
// Method must return multi-line text
// You must define which symbol will used as live cell
type DataSourceProvider interface {
	GetData() (string, error)
	GetLiveSymbol() string
}

// Implements loading initial input data from text file
type FileSourceDataProvider struct {
	FilePath string
}

func NewFileSourceDataProvider(filePath string) *FileSourceDataProvider {
	return &FileSourceDataProvider{FilePath: filePath}
}

func (r *FileSourceDataProvider) GetData() (string, error) {
	b, err := ioutil.ReadFile(r.FilePath)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Defines symbol that will be interpreted as LIVE cell
func (r *FileSourceDataProvider) GetLiveSymbol() (string) {
	return "O"
}

// Implements creating random initial input data
// You can use density variable to set how sparse/dense will be initial population

const DefaultPopulationDensity = 2

type RandomSourceDataProvider struct {
	Size int
	Density int
}

// Used it for generating random input data
// Size defines size of desk to be created (NxN where N = size)
func NewRandomSourceDataProvider(size int, density int) *RandomSourceDataProvider {
	if density <= 0 {
		density = DefaultPopulationDensity
	}
	return &RandomSourceDataProvider{Size: size, Density: density}
}

func (r *RandomSourceDataProvider) GetData() (string, error) {
	if r.Size <= 0 {
		return "", ErrInvalidDeskSize
	}

	data := ""

	for i := 0; i < r.Size; i++ {
		for j := 0; j < r.Size; j++ {
			if !r.GetRandomState() {
				data += "-"
			} else {
				data += r.GetLiveSymbol()
			}
		}

		data += "\n"
	}

	return data, nil
}

// Defines symbol that will be interpreted as LIVE cell
func (r *RandomSourceDataProvider) GetLiveSymbol() (string) {
	return "O"
}

// Generates random true/false
// Attention! It's a developer's responsible to provide non-zero density value
func (r *RandomSourceDataProvider) GetRandomState() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) % r.Density == 0
}