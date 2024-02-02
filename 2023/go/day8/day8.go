package main

/*
Riding a camel through the desert
Use map to navigate. 
    - List of left/right instructions can be repeated
    - node network from AAA to ZZZ
How many steps are required to get to ZZZ?

*/

import (
	"aoc/common"
	"fmt"
	"regexp"
)

func parseNodes(lines []string) map[string][]string {
    treenodes := make(map[string][]string)
    const tokenpattern = "[A-Z]+"
    re := regexp.MustCompile(tokenpattern)
    for _, line  := range lines {
        matches := re.FindAllString(line, -1)
        if len(matches) == 3 {
            treenodes[matches[0]] = matches[1:]
        }
    }
    return treenodes
}

func traverseToEnd(nodemapping map[string][]string, instructions string) int{
    steps := 0 
    currentNode := "AAA"
    for currentNode != "ZZZ" {
        for _, instrrune := range instructions {
            if instrrune == 'L' {
                currentNode = nodemapping[currentNode][0]
            } else if instrrune == 'R' {
                currentNode = nodemapping[currentNode][1]
            } else {
                fmt.Println("Bad run encountered: ", instrrune)
            }
            steps++
        }
    }
    return steps
}

func main() {
    lines := common.LoadFileLines("input.txt") 
    instructions := lines[0]
    fmt.Println(instructions)
    nodemapping := parseNodes(lines[2:])
    stepcount := traverseToEnd(nodemapping, instructions)
    fmt.Println(stepcount)
} 
