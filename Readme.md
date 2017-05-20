# Conway's Game of Life
GO implementation of "The Game of Life", also known simply as Life, devised by the British mathematician John Horton Conway in 1970.

You can read about the rules at [Wiki](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Rules)

Naive approach used with two arrays for storing current and next generation states.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for testing purposes.

### Prerequisites

To be able to run the game locally you should have

```
linux machine
git tool
docker or golang installed
```

### Quick start with Docker

```
cd ~
git clone https://github.com/spaiz/gogameoflife.git && cd gogameoflife
docker build --rm -t golife .
```

To play the game in RANDOM mode, run:

```
docker run -it --rm --name="game_of_life" golife -random=50 -density=4 -delay=100
```

To play the game in MANUAL mode, you must provide path to the file with input data:

```
docker run -v $PWD/data/:/data -it --rm --name="game_of_life" golife -file=/data/input.txt
```

## Running the tests

```
cd ./tests
go test -v
```

## Game modes

### Random mode:

You can start the game in "RANDOM" mode by providing an "-random" parameter as a positive non zero integer,
that will be used as desk size.

This command will run the game with generated 100x100 desk and initial data with live cells population density level 3.
New generation will be created every 500 milliseconds.

```
golife -random=100 -density=3 -delay=500
```

### Manual mode:

Text file may be used as initial data source. Please see test data in the file "./data/input.txt". An valid file is a text file, where the number of symbols in each line equals the number of lines. (NxN matrix). A live cells represented by big "O".

By default, application will look for "input.txt" file in the current working directory. 

Example:

```
cd ./data/
golife -delay=1000
```

Or with a custom file path:

```
golife -file=/path/to/your/file/input.txt -delay=1000
```

## Install
If you have GO installed locally, you can compile and run the game locally. All dependencies are already included in the repository.

```
cd ~
git clone https://github.com/spaiz/gogameoflife.git && cd gogameoflife
cd ./cmd/golife
go install
golife -random=100 -density=3 -delay=200
```

## Command Line parameters
Please, use the next command to see all available flags and their default values.

```
golife -h
```