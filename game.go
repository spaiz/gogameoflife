package gogameoflife

import (
	"time"
	"os"
	"os/exec"
	"fmt"
	"github.com/exrook/drawille-go"
	"io"
)

// Min delay
// Very naive approach.
// Current desc state used as initial state for each iteration
// Every iteration we print current desk, then create new array (slice) for storing the next state
// Then we set new state as current state

type Game struct {
	CurrentGenerationDesk Desk
	NextGenerationDesk    Desk
	InitialDataProvider   DataSourceProvider
	Canvas                drawille.Canvas
}

func NewGame(provider DataSourceProvider) *Game {
	return &Game{
		InitialDataProvider: provider,
		Canvas: drawille.NewCanvas(),
	}
}

// Delay should be provided in milliseconds
func (r *Game) Start(delay int) error {
	fileContent, err := r.InitialDataProvider.GetData()

	if err != nil {
		return err
	}

	r.CurrentGenerationDesk, err = ParseInputData(fileContent, r.InitialDataProvider.GetLiveSymbol())

	if err != nil {
		return err
	}

	for {
		r.ClearScreen()
		r.Render(os.Stdout)
		r.NextGeneration()
		r.Refresh()
		r.Sleep(delay)
	}

	return nil
}

// Creates new generation of population
// Naive and non-optimal solution. Every iteration new desk created. Hopefully GC will fine with it.
// If necessary, may be optimized to reuse 2 predefined arrays
// with switch() and clear() every iteration
func (r *Game) NextGeneration() Desk {
	r.NextGenerationDesk = CreateDefaultDeadDesk(len(r.CurrentGenerationDesk))

	for i := range r.CurrentGenerationDesk {
		for j := range r.CurrentGenerationDesk {
			r.NextGenerationDesk[i][j] = r.CurrentGenerationDesk.GetNewCellState(i, j)
		}
	}

	return r.NextGenerationDesk
}

// Set printed new generation as current generation
// NextGeneration() method will use it as initial data on next iteration
func (r *Game) Refresh() {
	r.CurrentGenerationDesk = r.NextGenerationDesk
}

// Clear the terminal (linux only)
func (r *Game) ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// Renders current desk to console
// Some package used to emulate pixels in terminal
func (r *Game) Render(out io.Writer) {
	r.Canvas.Clear()

	for i, x := range r.CurrentGenerationDesk {
		for j, v := range x {
			if v {
				r.Canvas.Set(j, i)
			}
		}
	}

	fmt.Fprintf(out, fmt.Sprintf("%s", r.Canvas))
}

// Used for delay between creating new generation
// Used for controlling "screen refresh rate"
func (r Game) Sleep(delay int) {
	time.Sleep(time.Millisecond * time.Duration(delay))
}