/*
Need to add up all the numbers to get the Elf's gondola lift working

Engine schematic consists of visual representation of the engine.

Any number adjacent to a symbol is a "part number" and should be included in the sum

# Periods do not count as symbols

Seems like columns and rows need to be accounted for
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/exp/slices"
)

type gearNumber struct {
    number int
    startCol int
    endCol int
}

func getRowColMax(inputLines[]string) (int, int){
    rowMax := len(inputLines) - 1
    colMax := len(inputLines[0]) - 1
    return rowMax, colMax
}


func CreateSymMap(inputLines []string) map[int][]int {
    symMap := make(map[int][]int)
    for rowNum, line := range inputLines {
        for colNum, rune := range line {
            if rune != '.' && !unicode.IsDigit(rune){
                symMap[rowNum] = append(symMap[rowNum], colNum)
            }
        }
    }
    return symMap
}

// A part number is next to special symbol.  This can be on the row above and below the number and diagonally touching
func IsPartNumber(row int, start int, end int, symMap map[int][]int, rowMax int, colMax int) bool{
    res := false
    startCheck := start -1 
    if start == 0 {
        startCheck = 0
    }
    endCheck := end + 1
    if endCheck > colMax{
        endCheck = colMax
    }

    //fmt.Println("Row: ", row, " sCol: ",startCheck, " eCol: ", endCheck, " Map: ", symMap[row])
    if symMap[row] != nil && 
       (slices.Contains(symMap[row], startCheck) || 
        slices.Contains(symMap[row], endCheck) ) {
        //fmt.Println("found a num in row")
        res = true 
    } else {
        rowsCheck := make([]int, 0)
        if row == 0 {
            rowsCheck = append(rowsCheck, row + 1)
        } else if row == rowMax {
            rowsCheck = append(rowsCheck, row - 1)
        } else {
            rowsCheck = append(rowsCheck, row + 1)
            rowsCheck = append(rowsCheck, row - 1)
        }

        for _, currRow := range rowsCheck{
            //fmt.Println("Map ", currRow, ": ", symMap[currRow])
            for i:= startCheck; i <= endCheck; i++ {
                if symMap[currRow] != nil &&
                   slices.Contains(symMap[currRow], i) {
                   res = true 
                    //fmt.Println("found a num in adj row")
                }
            }
        }
    }
    return res
}

func FindPartNumbers(inputLines[]string, symMap map[int][]int, rowMax int, colMax int) []int {
    partNumbers := make([]int, 0)
    for row, line := range inputLines {
        //fmt.Println("******************** Row: ", row , "********************")
        foundDigit := false
        colStart := 0
        colEnd := 0
        for col, rune := range line {
            if !foundDigit {
                if unicode.IsDigit(rune){
                    foundDigit = true 
                    colStart = col
                }
            } else {
                if !unicode.IsDigit(rune) ||
                   (unicode.IsDigit(rune) && col == colMax) {
                    foundDigit = false 
                    colEnd = col
                    if IsPartNumber(row, colStart, colEnd - 1 , symMap, rowMax, colMax) {
                        var num int
                        var err error
                        if col == colMax && unicode.IsDigit(rune){
                            num, err = strconv.Atoi(line[colStart:])
                        } else {
                            num, err = strconv.Atoi(line[colStart:colEnd])
                        }
                        if err != nil {
                            log.Fatal(err)
                        } else{
                            //fmt.Println("cs: ", colStart, " ce: ", colEnd)
                            //fmt.Println("found num: ", num)
                            partNumbers = append(partNumbers, num) 
                        }
                    }
                }
            }
        }
    }
    return partNumbers
}

// finds gears regardless if numbers are adjacent or not
func FindGears(inputLines []string) map[int][]int{
    symMap := make(map[int][]int)
    for rowNum, line := range inputLines { 
        for colNum, rune := range line {
            if rune == '*' {
                symMap[rowNum] = append(symMap[rowNum], colNum)
            }
        }
    }
    return symMap
}

func parseNum(rowString string, colStart int, colEnd int, colMax int) gearNumber{
    var numSlice string
    if colEnd == colMax + 1 {
        numSlice = rowString[colStart:]
    } else {
        numSlice = rowString[colStart:colEnd]
    }
    num, err := strconv.Atoi(numSlice)
    if err != nil{
        log.Fatal(err)
    }
    if colEnd != colMax {
        colEnd = colEnd - 1
    }
    number := gearNumber{
        number : num,
        startCol : colStart,
        endCol: colEnd,
    }
    return number
}

func findNumbers(rowString string, colMax int) []gearNumber{
    var numbers  []gearNumber
    foundNum := false
    colStart := 0
    for colNum, rune := range rowString {
        if foundNum == false{
            if unicode.IsDigit(rune) {
                foundNum = true
                colStart = colNum
                if colNum == colMax {
                    num := parseNum(rowString, colStart, colNum, colMax)
                    numbers = append(numbers, num)
                    foundNum = false
                }
            }
        } else {
            if !unicode.IsDigit(rune) ||
               (colNum == colMax && unicode.IsDigit(rune)) {
                colEnd := colNum
                if unicode.IsDigit(rune) && colNum == colMax{
                    fmt.Println("found :", string(rune), " at  col: ", colNum)
                    colEnd = colMax + 1
                }
                fmt.Println("Col start: ", colStart)
                fmt.Println("Col end: ", colEnd)
                fmt.Println("Col Max: ", colMax)
                fmt.Println("Substring: ", rowString[colStart:colEnd])
                num := parseNum(rowString, colStart, colEnd, colMax) 
                numbers = append(numbers, num)
                foundNum = false
            }
        }
    }
    return numbers
}

func isNumbersInAdjCol(number gearNumber, col int, colMax int) bool {
    result := false    
    colUpperCheck := col +1 
    if colUpperCheck > colMax {
       colUpperCheck = colMax 
    }
    colLowerCheck := col - 1
    if colLowerCheck < 0 {
       colLowerCheck = 0 
    }
    if (number.startCol >= colLowerCheck && number.startCol <= colUpperCheck) ||
       (number.endCol >= colLowerCheck && number.endCol <= colUpperCheck) {
       result = true 
    }
    return result
}

// find gear partners where gears have exactly 2 adjacent numbers
func FindGearPartners(inputLines []string, symMap map[int][]int, rowMax int, colMax int) [][]int{
    var gearPairs [][]int
    for rowNum, colNums := range symMap {
        checkRows := []int{rowNum}
        if rowNum - 1 >= 0 {
          checkRows = append(checkRows, rowNum - 1) 
        }
        if rowNum + 1 <= rowMax {
          checkRows = append(checkRows, rowNum + 1) 
        }
        for _, colNum := range colNums {
            var gearNums []int
            for _, row := range checkRows {
                numbers := findNumbers(inputLines[row], colMax)
                //fmt.Println("Gear col: ",colNum)
                //fmt.Println(numbers)
                for _, number := range numbers {
                    if isNumbersInAdjCol(number, colNum, colMax) {
                       gearNums = append(gearNums, number.number) 
                    } 
                }
            }
            if len(gearNums) == 2 {
                gearPairs = append(gearPairs, gearNums)
                //fmt.Println(rowNum, gearNums)
            } else{
                //fmt.Println("Bad Gear at row: ", rowNum + 1, " col: ", colNum + 1)
                //for _, num := range gearNums {
                    //fmt.Println("Found number: ", num)
                //}
            }
        }
    }
    return gearPairs
}

func Part1(inputLines []string) int{
    rowMax, colMax := getRowColMax(inputLines)
    symMap := CreateSymMap(inputLines)
    partNums := FindPartNumbers(inputLines, symMap, rowMax, colMax)
    //fmt.Println(partNums)
    sum := 0
    for _, num := range partNums {
        sum += num
    }

    fmt.Println(sum)
    return sum
}

func Part2(inputLines []string) int{
    rowMax, colMax := getRowColMax(inputLines)
    fmt.Println("Row Max: ", rowMax, " Col Max: ", colMax)
    // find gear and the adjacent part numbers
    gears := FindGears(inputLines)
    gearPairs := FindGearPartners(inputLines, gears, rowMax, colMax)
    //fmt.Println("Gear pairs: ", gearPairs)
    sum := 0
    for _, pairs := range gearPairs{
        sum += pairs[0] * pairs[1]
    }

    fmt.Println("Sum: ", sum)
    return sum
}

func main(){
    byteData, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    } 

    input := string(byteData)
    inputLines := strings.Split(input, "\n")
    inputLines = inputLines[0:len(inputLines)-1]

    //Part1(inputLines)
    Part2(inputLines)
}
