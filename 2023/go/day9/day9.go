package main

/*
Oasis and Sand Instability sensor
Goal: predict next value for each value provided and it's history

Find sequence from the difference at each step of history
recursively get the difference at each step of the sequence until you get a sequence of zeros

Then extrapolate the next value by reversing the process.
Add a zero to the sequence of zeros and reverse the calculation

sum of extrapolated values
*/

import (
	"aoc/common"
	"fmt"
	"regexp"
	"strconv"
)

const numpattern = "-?[0-9]+"

func parseSequences(lines []string) [][]int{
    var sequences [][]int
    re := regexp.MustCompile(numpattern)
    for _, line := range lines {
        matches := re.FindAllString(line, -1)
        var seq []int
        for _, match := range matches {
            num, err := strconv.Atoi(match)
            if err != nil {
                fmt.Println(err)
            }
            seq = append(seq, num)
        }
        if len(seq) >0 {
            sequences = append(sequences, seq)
        }
    }
    return sequences
}

func isZeroSeq(seq []int) bool {
    zeroseq := true
    for _, num := range seq {
        if num != 0 {
            zeroseq = false
            break
        }
    }
    return zeroseq
}

func getNextDiffSeq(seq []int) []int{
    var diffSeq []int
    for index:= 1; index < len(seq); index++ {
        diff := seq[index] - seq[index-1]
        diffSeq = append(diffSeq, diff)
    }
    return diffSeq
}

func extrapolateSeq(sequences [][]int) int {
    extraval := 0
    for index:= len(sequences)-1; index >= 0; index--{
        extraval += sequences[index][len(sequences[index])-1]
    }
    return extraval
}

func getExtrapolatedValues(sequences [][]int) []int {
    var extrapolatedVals []int
    for _, seq := range sequences {
        currentSeq := seq
        var diffseqs [][]int
        diffseqs = append(diffseqs, currentSeq)

        for isZeroSeq(currentSeq) == false {
            diffseq := getNextDiffSeq(currentSeq)
            diffseqs = append(diffseqs, diffseq)
            currentSeq = diffseq
        }

        extraval := extrapolateSeq(diffseqs)
        extrapolatedVals = append(extrapolatedVals, extraval)
    }
    return extrapolatedVals
}

func main() {
    lines := common.LoadFileLines("input.txt") 
    sequences := parseSequences(lines)
    extrapolatedVals := getExtrapolatedValues(sequences)

    sum := 0
    for _, val := range extrapolatedVals { 
        sum += val
    }
    fmt.Println(sum)
}
