package route

import (
	"errors"
	"sync"
)

type Direction int

const (
	A                 = 'A'
	B                 = 'B'
	Unknown Direction = iota
	Left
	Right
	Up
	Down
)

type Coord struct {
	column, row int
}

type CellUpdate struct {
	Column, Row int
	Value       rune
}

type Navigator struct {
	mapSize, coordsA, coordsB Coord
	aIsPrimary                bool
	directionA, directionB    Direction
	resultMap                 [][]rune
}

func NewNavigator(rowCount int, colCount int) *Navigator {
	return &Navigator{
		mapSize:   Coord{colCount, rowCount},
		resultMap: make([][]rune, rowCount),
	}
}

func (n *Navigator) ParseLine(line string, row int) error {
	n.resultMap[row] = make([]rune, n.mapSize.column)
	for col, symbol := range line {
		n.resultMap[row][col] = symbol
		if symbol == A {
			n.coordsA = Coord{col, row}
			continue
		}
		if symbol == B {
			n.coordsB = Coord{col, row}
			continue
		}
	}
	return nil
}

func (n *Navigator) calculatePrimaryPoint() error {
	if n.coordsA == n.coordsB {
		return errors.New("dots are equal")
	}

	n.aIsPrimary = n.coordsA.row < n.coordsB.row ||
		(n.coordsA.row == n.coordsB.row && n.coordsA.column < n.coordsB.column)

	n.calculateDirection()
	return nil
}

func (n *Navigator) calculateDirection() error {
	if n.aIsPrimary {
		if n.coordsA.column%2 == 0 {
			n.directionA = Up
		} else {
			n.directionA = Left
		}
		if n.coordsB.column%2 == 0 {
			n.directionB = Down
		} else {
			n.directionB = Right
		}
	} else {
		if n.coordsA.column%2 == 0 {
			n.directionA = Down
		} else {
			n.directionA = Right
		}
		if n.coordsB.column%2 == 0 {
			n.directionB = Up
		} else {
			n.directionB = Left
		}
	}
	return nil
}

func (n *Navigator) FindRoute() error {
	var waitGroup sync.WaitGroup
	var waitGroupDraw sync.WaitGroup
	updates := make(chan CellUpdate, 256)

	draw := func(updates <-chan CellUpdate) {
		defer waitGroupDraw.Done()
		for u := range updates {
			cell := &n.resultMap[u.Row][u.Column]
			if *cell != 'A' && *cell != 'B' && *cell != '#' {
				*cell = u.Value
			}
		}
	}

	findRoute := func(forA bool, updates chan<- CellUpdate) {
		defer waitGroup.Done()
		dotToDraw := 'b'
		if forA {
			dotToDraw = 'a'
		}
		goUp := func(dot *Coord) {
			for i := dot.row - 1; i >= 0; i-- {
				updates <- CellUpdate{dot.column, i, dotToDraw}
			}
			dot.row = 0
		}
		goDown := func(dot *Coord) {
			for i := dot.row + 1; i < n.mapSize.row; i++ {
				updates <- CellUpdate{dot.column, i, dotToDraw}
			}
			dot.row = n.mapSize.row - 1
		}
		goLeft := func(dot *Coord) {
			for i := dot.column - 1; i >= 0; i-- {
				updates <- CellUpdate{i, dot.row, dotToDraw}
			}
			dot.column = 0
		}
		goRight := func(dot *Coord) {
			for i := dot.column + 1; i < n.mapSize.column; i++ {
				updates <- CellUpdate{i, dot.row, dotToDraw}
			}
			dot.column = n.mapSize.column - 1
		}

		if n.aIsPrimary {
			if forA {
				if n.directionA == Up {
					goUp(&n.coordsA)
					goLeft(&n.coordsA)
					return
				}
				if n.directionA == Left {
					goLeft(&n.coordsA)
					goUp(&n.coordsA)
					return
				}
			} else {
				if n.directionB == Down {
					goDown(&n.coordsB)
					goRight(&n.coordsB)
					return
				}
				if n.directionB == Right {
					goRight(&n.coordsB)
					goDown(&n.coordsB)
					return
				}
			}
		} else {
			if forA {
				if n.directionA == Down {
					goDown(&n.coordsA)
					goRight(&n.coordsA)
					return
				}
				if n.directionA == Right {
					goRight(&n.coordsA)
					goDown(&n.coordsA)
					return
				}
			} else {
				if n.directionB == Up {
					goUp(&n.coordsB)
					goLeft(&n.coordsB)
					return
				}
				if n.directionB == Left {
					goLeft(&n.coordsB)
					goUp(&n.coordsB)
					return
				}
			}
		}
	}
	n.calculatePrimaryPoint()

	waitGroup.Add(2)
	go findRoute(true, updates)
	go findRoute(false, updates)

	waitGroupDraw.Add(1)
	go draw(updates)
	waitGroup.Wait()
	close(updates)
	waitGroupDraw.Wait()
	return nil
}

func (n *Navigator) GetResult() []string {
	result := make([]string, n.mapSize.row)

	for i, runes := range n.resultMap {
		result[i] = string(runes)
	}

	return result
}
