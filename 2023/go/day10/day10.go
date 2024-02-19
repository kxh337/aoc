package main

/*
Pipe Maze

Pipe rules:
    | is a vertical pipe connecting north and south.
    - is a horizontal pipe connecting east and west.
    L is a 90-degree bend connecting north and east.
    J is a 90-degree bend connecting north and west.
    7 is a 90-degree bend connecting south and west.
    F is a 90-degree bend connecting south and east.
    . is ground; there is no pipe in this tile.
    S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

Assume pipe is one large continuous loop but some pipes are not connected to the loop

Determine the # of steps in the loop from the starting position to the point farthest from the start position
*/

import (
	"aoc/common"
	"fmt"
)

type PipeDir int

const (
   NORTH PipeDir = iota
   SOUTH
   EAST
   WEST
)

var PIPE_CHARS = map[rune][]PipeDir {
    '|' : {NORTH, SOUTH},
    '-' : {EAST, WEST},
    'L' : {NORTH, EAST},
    'J' : {NORTH, WEST},
    '7' : {SOUTH, WEST},
    'F' : {SOUTH, EAST},
    'S': {NORTH, SOUTH, EAST, WEST},
}

func covertToRuneTable(rows []string) [][]rune{
    var runeTable [][]rune
    for _, row := range rows {
        var rowRunes []rune
        for _, rune := range row {
           rowRunes = append(rowRunes, rune)
        }
        runeTable = append(runeTable, rowRunes)
    }
   return runeTable
}

func pipeExists(row int, col int, mapData map[int][]int) bool{
    res := false
    if mapData[row] != nil {
        for _, pipe := range mapData[row]{
            if pipe == col {
                res = true
            }
        }
    }
    return res
}

// @TODO:maybe need to check for bounds? seems like the pipe never leads out of bounds so we're likely okay
func findPipeNetwork(row int, col int, mapData map[int][]int, runeTable [][]rune) map[int][]int{
    curRune := runeTable[row][col]
    options, ok := PIPE_CHARS[curRune]
    if ok {
        for _, option := range options {
            checkrow := row
            checkcol := col
            var checkOption PipeDir
            switch option {
                case NORTH:
                    checkrow--
                    checkOption = SOUTH
                case SOUTH:
                    checkrow++
                    checkOption = NORTH
                case EAST:
                    checkcol++
                    checkOption = WEST
                case WEST:
                    checkcol--
                    checkOption = EAST
            }

            if !pipeExists(checkrow, checkcol, mapData) {
                //fmt.Println("Checking: ", checkrow, checkcol)
                checkOptions := PIPE_CHARS[runeTable[checkrow][checkcol]]
                //fmt.Println("Looking for: ", checkOption)
                //fmt.Println("Found options: ", checkOptions)
                for _, option := range checkOptions {
                    if option == checkOption{
                        //fmt.Println("Found neighbor")
                        if mapData[checkrow] == nil{
                            mapData[checkrow] = make([]int, 0)
                        }
                        mapData[checkrow] = append(mapData[checkrow], checkcol)
                        return findPipeNetwork(checkrow, checkcol, mapData, runeTable)
                    }
                }
            }
        }
    } else {
        // Assume no options and return
        return mapData
    }
    return mapData
}

func getBreadthNextNeighbors(currentPipe map[int][]int, visitedNodes map[int][]int, mapData map[int][]int) map[int][]int{
    nextPipes := map[int][]int{}
    for curRow, curCols := range currentPipe {
        for _, curCol := range curCols {
            checkRows := []int{curRow}
            if  curRow - 1 >= 0 {
                checkRows = append(checkRows, curRow - 1)
            }
            if  curRow + 1 <= 139 {
                checkRows = append(checkRows, curRow + 1)
            }
            checkCols := []int{curCol}
            if  curCol - 1 >= 0 {
                checkCols = append(checkCols, curCol - 1)
            }
            if  curCol + 1 <= 139 {
                checkCols = append(checkCols, curCol + 1)
            }

            for _, checkRow := range checkRows {
                for _, checkCol := range checkCols {
                    // fmt.Println(checkRow, checkCol)
                    addPipe := true
                    // check that it isn't visted first
                    visitedRow, ok := visitedNodes[checkRow]
                    if ok {
                        for _, col := range visitedRow {
                            if col == checkCol {
                                // fmt.Println("Pipe visited")
                                addPipe = false
                                break
                            }
                        }
                    }
                    // check that it is in the map
                    if addPipe {
                        // fmt.Println("checking if pipe is on map")
                        // assume that map does not have the pipe
                        addPipe = false
                        mapRow, ok := mapData[checkRow]
                        if ok {
                            for _, col := range mapRow {
                                if col == checkCol{
                                    // fmt.Println("Found it on the map")
                                    addPipe = true
                                    break
                                }
                            }
                        }
                    }
                    if addPipe {
                        if nextPipes[checkRow] == nil {
                            nextPipes[checkRow] = make([]int, 0)
                        }
                        nextPipes[checkRow] = append(nextPipes[checkRow], checkCol)
                    }
                }
            }
        }
    }

    // append  nextPipes to visited nodes
    for row,  cols := range nextPipes{
        visitedNodes[row] = append(visitedNodes[row], cols...)
    }
    return nextPipes
}

// there is a loop in the pipes, i.e. pipes only connect to 2 other pipes
func getLongestPathTile(startRow int, startCol int, mapData map[int][]int) int{
    visitedPipes := map[int][]int{startRow: {startCol}}
    startPipe := map[int][]int{startRow: {startCol}}
    nextPipes := getBreadthNextNeighbors(startPipe, visitedPipes, mapData)
    stepCount := 1

    for (len(nextPipes) > 0) {
        fmt.Println(len(nextPipes))
        nextPipes = getBreadthNextNeighbors(nextPipes, visitedPipes, mapData)
        stepCount++
    }
    return stepCount
}

func main() {
    rows := common.LoadFileLines("input.txt")
    runeTable := covertToRuneTable(rows)
    foundStart := false
    mapData := make(map[int][]int)
    startCol := -1
    startRow := -1
    for rowIndex, row := range rows {
        for colIndex, curRune:= range row {
            if curRune == 'S' {
                mapData[rowIndex] = []int{colIndex}
                startRow = rowIndex
                startCol = colIndex
                foundStart = true
                break
            }
        }
        if foundStart {
            break
        }
    }
    if foundStart {
        pipemap := findPipeNetwork(startRow, startCol, mapData, runeTable)
        pipeCount := 0
        for _, row := range pipemap {
            pipeCount += len(row)
        }
        fmt.Println(pipemap)
        fmt.Println(pipeCount/2)
        //fmt.Println("Step  count: ", getLongestPathTile(startRow, startCol, pipemap))
    }else {
        fmt.Println("Failed to find the starting pipe")
        return
    }

}
