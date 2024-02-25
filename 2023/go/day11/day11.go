package main

import (
	"aoc/common"
	"math"
	"fmt"
	"slices"
)


func addRow(lines []string, rowIndex int, colSize int) []string{
    hashRow := ""
    for idx := 0; idx < colSize; idx ++ {
        hashRow  += "."
    }

    lines = slices.Insert(lines, rowIndex, hashRow)

    return lines
}

func addCol(lines[]string, colIndex int, rowSize  int) []string{
    for index, line := range lines {
        lines[index] = line[:colIndex] + "." + line[colIndex:]
    }
    return lines
}

type Galaxy struct {
    id int
    rowInd int
    colInd int
}

func main() {
    lines := common.LoadFileLines("input.txt")
    MAX_ROW := len(lines)
    MAX_COL := len(lines[0])

    var rowsToAdd []int
    var colsWithHash []int
    var galaxyMap []Galaxy

    // keep track of lines and cols without # char
    for rowInd, line := range lines {
        hashFound := false
        for colInd, rune := range line {
            if rune == '#' {
                hashFound =true
                if !slices.Contains(colsWithHash, colInd){
                    colsWithHash = append(colsWithHash, colInd)
                }
            }
        }
        if !hashFound{
           rowsToAdd = append(rowsToAdd, rowInd)
        }

    }

    fmt.Println("Galaxy expansion")
   // insert the rows and columns
   slices.Reverse(rowsToAdd)
    for _, rowInd := range rowsToAdd {
        //fmt.Println("Adding row at ", rowInd)
        lines = addRow(lines, rowInd, MAX_COL)
        MAX_ROW++
    }

    // Add column if column does not have a hash
    for colInd := MAX_COL - 1; colInd >= 0 ; colInd-- {
        if !slices.Contains(colsWithHash, colInd){
            //fmt.Println("Adding col at ", colInd)
            lines = addCol(lines, colInd, MAX_ROW)
            MAX_COL++
        }
    }

    fmt.Println("Finding Galaxies")
    id := 0
    // remap the galaxies after expansion
    for rowInd, line := range lines {
        for colInd, rune := range line {
            if rune == '#' {
                gal := Galaxy{id, rowInd, colInd}
                galaxyMap = append(galaxyMap, gal)
                //fmt.Println(gal)
                id++
            }
        }
    }
    //fmt.Println("Found ", id, " many hash tags")

    fmt.Println("Calculating distances")
    pairs := make(map[int][]int)
    var lengths []int
    // find distances to each pair
    for _, galaxy := range galaxyMap{
        for _, pot_pair := range galaxyMap{
            if pot_pair != galaxy {
                res1, exist1 := pairs[galaxy.id]
                res2, exist2 := pairs[pot_pair.id]

                // if pair doesn't already exist then add it
                if (!exist1 || !slices.Contains(res1, pot_pair.id)) &&
                   (!exist2 || !slices.Contains(res2, galaxy.id)) {
                    pairs[galaxy.id] = append(pairs[galaxy.id], pot_pair.id)
                    pair_dst := int(math.Abs(float64(galaxy.rowInd) - float64(pot_pair.rowInd)) +
                                    math.Abs(float64(galaxy.colInd) - float64(pot_pair.colInd)))
                    lengths = append(lengths, pair_dst)
                    //fmt.Println("New pair", galaxy, pot_pair, " Dist: ", pair_dst)
                }
            }
        }
    }
    fmt.Println("Found :", len(lengths), " pairs")

    var sum int
    sum = 0
    for _, length := range lengths {
        sum = sum + length
    }
    fmt.Println(sum)
}
