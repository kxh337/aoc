/*
Cube conundrum

Cubes are red, green or blue
Figure out information about the number of cubes
Elf reaches into the bag and grabs a handful of random cubes, show them to you and then put them back into the bag.
Games can have multiple sets which are separated by semicolons

Elf asks which games were possible.  Then add up the sum of the IDs of the possible games

Part 2
what is the fewest number of cubes of each color that could have been the bag to make the game possible
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

// Defined by the elf lol
var CubeLimit = map[string]int {
    "red" : 12,
    "green" : 13,
    "blue" : 14,
}

type Game struct {
    ID int
    sets []map[string]int
}

func ParseInput(lines []string) []Game {
    var gameSets []Game
    const RENUM string = "([0-9])+"
    regexNum := re.MustCompile(RENUM)
    const RECOLOR string = "(red|green|blue)"
    regexColor := re.MustCompile(RECOLOR)
    for _, line := range lines {
        game := Game{}
        params := strings.Split(line, ":")
        num_match := regexNum.FindAllString(params[0], -1)
        game_id, err := strconv.Atoi(num_match[0])
        if err != nil {
            log.Fatal(err)
        }
        game.ID = game_id

        sets := strings.Split(params[1], ";")
        for setNum, set := range sets {
            game.sets = append(game.sets, make(map[string]int))
            marbles := strings.Split(set, ",")
            for _, marble := range marbles{
                countString := regexNum.FindAllString(marble, -1)
                color := regexColor.FindAllString(marble, -1)
                count, err := strconv.Atoi(countString[0])
                if err != nil {
                    log.Fatal(err)
                }
                game.sets[setNum][color[0]] = count
            }
        }
        gameSets = append(gameSets, game)
    }
    return gameSets
}

func FilterPossibleGames(gameSets []Game) []Game {
    var goodGames []Game 
    for _, game := range gameSets {
        isGoodGame := true
        for _, set := range game.sets {
            if CubeLimit["red"] <  set["red"] ||
               CubeLimit["green"] <  set["green"] ||
               CubeLimit["blue"] <  set["blue"] {
                   isGoodGame = false
                   break
            }
        }
        if isGoodGame {
            goodGames = append(goodGames, game)
        }
    }
    return goodGames
}

func sol1(gameData []Game){
    goodGames := FilterPossibleGames(gameData)
    fmt.Println("Good Games: ", goodGames)
    sum := 0
    for _, game := range goodGames {
        fmt.Println(game.ID)
        sum += game.ID
    }
    fmt.Println("Sum game IDs: ", sum)

}

func sol2(gameData []Game){
    var minCubeMaps []map[string]int

    for game_idx, game := range gameData {
        minCubeMaps = append(minCubeMaps, make(map[string]int))
        minCubes := minCubeMaps[game_idx]
        minCubes["red"] = 0
        minCubes["blue"] = 0
        minCubes["green"] = 0
        for _, set := range game.sets {
            if set["red"] > minCubes["red"] {
                minCubes["red"] = set["red"]
            }
            if set["blue"] > minCubes["blue"] {
                minCubes["blue"] = set["blue"]
            }
            if set["green"] > minCubes["green"] {
                minCubes["green"] = set["green"]
            }
        }
    }

    sum := 0
    for _, minCubeMap := range minCubeMaps {
        fmt.Println(minCubeMap)
        sum += minCubeMap["red"] * minCubeMap["blue"] * minCubeMap["green"]
    }
    fmt.Println(sum)
}

func main(){
    byteData, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    } 
    input := string(byteData)
    input_lines := strings.Split(input, "\n")
    gameData := ParseInput(input_lines[0:len(input_lines)-1])

    //sol1(gameData)
    sol2(gameData)
}
