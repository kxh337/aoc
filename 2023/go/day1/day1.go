/*
Trebuchet Calibration

Each line had a calibration value that the Elves need to recover.
Each line, the calibration value can be found by combining the 1st and last digit to form a single two-digit number
*/
package main

import (
	"fmt"
	"log"
	"os"
	re "regexp"
	"strconv"
	"strings"
)

var NumberMap = map[string]string {
    "one"   : "1",
    "two"   : "2",
    "three" : "3",
    "four"  : "4",
    "five"  : "5",
    "six"   : "6",
    "seven" : "7",
    "eight" : "8",
    "nine"  : "9",
}

func ReplaceNumberStrings(lines []string) []string {
    var resultLines []string
    for _, line := range lines {
       var strBuilder strings.Builder
       for _, char := range line {
           strBuilder.WriteRune(char)
           for numberStr, number := range NumberMap {
               if strings.Contains(strBuilder.String(), numberStr) {
                   line  = strings.Replace(line, numberStr, number + string(char), 1)
                   fmt.Println(line, number)
                   strBuilder.Reset()
                   strBuilder.WriteRune(char)
                   break
               }
           }
       }
       fmt.Println(line)
       resultLines = append(resultLines, line)
   }
   return resultLines
}

func SumOfLines(lines []string) int {
    const REGEXPRSTR string = "([0-9]){1}"
    regexpr := re.MustCompile(REGEXPRSTR)
    sum := 0
    for _, subString := range lines {
        matches := regexpr.FindAllString(subString, -1)
        num1, err1 := strconv.Atoi(matches[0]) 
        num2, err2 := strconv.Atoi(matches[len(matches)-1])
        if err1 != nil || err2 != nil {
            log.Fatal(err1, err2)
        }
        fmt.Println(subString, num1, num2)
        cal := (num1 * 10) + num2
        sum += cal
    }
    return sum
}

func main(){
    var lines []string
    var subString string

    byteData, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    } 
    input := string(byteData)

    for _, char := range input {
        if char == '\n' {
            lines = append(lines, subString)
            subString = ""
        } else {
            subString += string(char)
        }
    }

    lines = ReplaceNumberStrings(lines)
    sum := SumOfLines(lines)
    fmt.Println(sum)
}
