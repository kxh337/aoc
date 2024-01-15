package main

import (
	"aoc/common"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

/*
Boat races
Part 1
Input: list of times allowed for reach race and the best distance ever recorded
Need to make sure you go father in each race than the current record holder
Charge and release to allow the boat to move, charge time counts against race time

*/

const NUMPATTERN = "(\\d+)"
const TIMEROW = 0
const DISTROW = 1

func grabnums(line string) []int{
    var nums [] int
    re := regexp.MustCompile(NUMPATTERN)
    matches := re.FindAllString(line, -1)
    for _, match := range matches {
        num, err := strconv.Atoi(match)
        if err != nil {
            log.Fatal(err)
        }
        nums = append(nums, num)
    }
    return nums
}

func parseRaceData(lines []string) ([]int, []int){
    times := grabnums(lines[TIMEROW])
    distances := grabnums(lines[DISTROW])
    return times, distances
}

func calDistance(holdtime int, racetime int) int {
    speed :=  holdtime
    timeleft := racetime - holdtime
    distance := speed * timeleft
    return distance
}

func part1() {
    lines := common.LoadFileLines("input.txt") 
    times, distances := parseRaceData(lines)
    var winlist [] int
    for index, racetime := range times {
        winningscenarios := 0
        for holdtime := 1; holdtime < racetime; holdtime++ {
            if (calDistance(holdtime, racetime) > distances[index]){
                winningscenarios++
            }
        }
        winlist = append(winlist, winningscenarios)
    }
    solution := 1
    for _, wincount := range winlist {
        solution = solution * wincount
    }
    fmt.Println(solution)
}

func main(){
    part1()
    //part2()
}
