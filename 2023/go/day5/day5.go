/*
Need to fined the lowest number that corresponds to any of the initial seed numbers

*/

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const seedPattern = "seeds:"
const seed2SoildPattern = "seed-to-soil map:"
const soil2FertPattern = "soil-to-fertilizer map:"
const fert2WaterPattern = "fertilizer-to-water map:"
const water2LightPattern = "water-to-light map:"
const light2TempPattern = "light-to-temperature map:"
const temp2HumidPattern = "temperature-to-humidity map:"
const humid2LocationPattern = "humidity-to-location map:"
const numPattern = "(\\s*\\d+)+"
const numOnlyPattern = "(\\d+)+"

type itemMapping struct {
    dst int
    src int
    mapRange int
}

func strSlice2IntSlice(input []string) []int{
    var numSlice []int
    for _, numStr := range input {
        num, err := strconv.Atoi(numStr)
        if err != nil {
            log.Fatal(err)
        }
        numSlice = append(numSlice, num)
    }
    return numSlice
}

func getSeeds(input string)[]int{
    re := regexp.MustCompile(seedPattern + numPattern)
    seedSecMatch := re.FindAllString(input, -1)
    reNum := regexp.MustCompile(numOnlyPattern)
    seedStrs := reNum.FindAllString(seedSecMatch[0], -1)
    return strSlice2IntSlice(seedStrs)
}

func getMap(input string, sectionPattern string) []itemMapping {
    var sectionMap []itemMapping
    re := regexp.MustCompile(sectionPattern + numPattern)
    sectionstr := re.FindAllString(input, -1)
    reNum := regexp.MustCompile(numOnlyPattern)
    sectionNumStrs := reNum.FindAllString(sectionstr[0], -1)
    sectionNums := strSlice2IntSlice(sectionNumStrs)
    
    for index, _ := range sectionNums{
        if (index + 1) % 3 == 0{
            dst := sectionNums[index - 2]
            src := sectionNums[index - 1]
            mapRange := sectionNums[index]
            itemMap := itemMapping {dst: dst, src: src, mapRange: mapRange}
            sectionMap = append(sectionMap, itemMap)
        }
    }

    return sectionMap
}

func getSeedLocations(seeds []int, mapping [][]itemMapping) []int{
    var locations []int
    for _, seed := range seeds {
        item := seed
        fmt.Println("")
        fmt.Println("Seed: ", item)
        for i, section := range mapping {
            fmt.Println("Finding mapping for: ", item)
            for _, itemMap := range section {
                fmt.Println("Checking mapping: dst: ", itemMap.dst, " src: ", itemMap.src,  " range: ",itemMap.mapRange)
                if item > itemMap.src && 
                   item < (itemMap.src + itemMap.mapRange) {
                    item =  itemMap.dst + (item - itemMap.src)
                    fmt.Println("Found in mapping #", i, ": ",itemMap, " Next Value: ", item)
                    break
                }

            }
        }
        locations = append(locations, item)
    }
    return locations
}
func part1(input string) {
    seeds := getSeeds(input)
    seed2SoilMap := getMap(input, seed2SoildPattern)
    soil2FertMap := getMap(input, soil2FertPattern)
    fert2WaterMap := getMap(input, fert2WaterPattern)
    water2LigtMap := getMap(input, water2LightPattern)
    light2TempMap := getMap(input, light2TempPattern)
    temp2HumidMap := getMap(input, temp2HumidPattern)
    humid2LocationMap := getMap(input, humid2LocationPattern)

    mappings := [][]itemMapping {seed2SoilMap, soil2FertMap, fert2WaterMap, water2LigtMap, light2TempMap, temp2HumidMap, humid2LocationMap}
    seedLocations := getSeedLocations(seeds, mappings)
    fmt.Println("seedLocations: ", seedLocations)
    
    closeestLocation := seedLocations[0]
    closeestLocationIndex := 0
    for index, location := range seedLocations{
        if location < closeestLocation{
            closeestLocation = location
            closeestLocationIndex = index
        }
    }
    fmt.Println("closest location index: ", closeestLocationIndex, " closest location value: ", closeestLocation)
}

func getSeedLIst(input string) map[int]int{
    seedInput := getSeeds(input)
    seedList := make (map[int]int)
    for i, num := range seedInput {
        if (i+1) % 2 == 0 {
            seedList[seedInput[i-1]] = seedInput[i-1] + num
        }
    }
    return seedList
}

func getItemMapVal(currVal int, itemMap itemMapping) int {
    res := -1
    if currVal >= itemMap.src && 
       currVal <= (itemMap.src + itemMap.mapRange) {
        res = (itemMap.src - currVal) + itemMap.dst
    }
    if res == -1 {
        fmt.Println("Error in getting the map value: ", currVal, " in map: ", itemMap)
    }
    return res
}

func updateUncoveredRange(newMin int, newMax int, uncoveredRange map[int]int) map[int]int {
    for currMin, currMax := range uncoveredRange {
       if (newMin >= currMin && newMin <= currMax) ||
          (newMax >= currMin && newMax <= currMax) {
            if newMin > currMin && newMax < currMax {
                // need to create two new sections
                uncoveredRange[newMax+1] = currMax
                uncoveredRange[currMin] = newMin - 1
            } else if newMin == currMin {
                uncoveredRange[newMax+1] = currMax
            } else if newMax == currMax {
                uncoveredRange[currMin] = newMin - 1
            }
            delete(uncoveredRange, currMin)
            // Should be okay to break as regions should not overlap when this is called
            break
       }
    }
    return uncoveredRange
}


func getNextRangeMap(currRangeMap map[int]int, itemMaps []itemMapping) map[int]int {
    nextRangeMap := make (map[int]int)
    for currMin, currMax := range currRangeMap {
        uncoveredRange := map[int]int {currMin : currMax}

        for _, itemMap := range itemMaps {
            var newMin int
            var newMax int
            //fmt.Println("currMin: ", currMin, " currMax: ", currMax, " itemMap: ", itemMap)

            if currMin <= itemMap.src {
                if currMax >= itemMap.src {
                    newMin = getItemMapVal(itemMap.src, itemMap)
                    if  currMax >= (itemMap.src + itemMap.mapRange){
                        fmt.Println("Min is below src and Max is above range")
                        newMax = getItemMapVal((itemMap.src + itemMap.mapRange), itemMap)
                    } else { // currMax < itemMaps.src + itemMaps.mapRange
                        fmt.Println("Min is below src and Max is below range")
                        newMax = getItemMapVal(currMax, itemMap)
                    }

                    nextRangeMap[newMin] = newMax
                    fmt.Println("newMax: ", newMax, " newMin: ",newMin)
                    uncoveredRange = updateUncoveredRange(newMin, newMax, uncoveredRange)
                }

            }else { // currMin > itemMaps.src
                if currMin <= itemMap.src + itemMap.mapRange {
                    newMin = getItemMapVal(currMin, itemMap)
                    if currMax <= itemMap.src + itemMap.mapRange {
                        fmt.Println("Min is above src and Max is below range")
                        newMax = getItemMapVal(currMax, itemMap)
                    } else {
                        fmt.Println("Min is above src and Max is above range")
                        newMax = getItemMapVal((itemMap.src + itemMap.mapRange), itemMap)
                    }
                    nextRangeMap[newMin] = newMax
                    fmt.Println("newMax: ", newMax, " newMin: ",newMin)
                    uncoveredRange = updateUncoveredRange(newMin, newMax, uncoveredRange)
                }
            }
        }

        for uncovMin, uncovMax := range uncoveredRange {
            // one to one mapping for uncoveredRanges
            nextRangeMap[uncovMin] = uncovMax
        }
    }
    return nextRangeMap
}


func getSeedListLocations(seedList map[int]int, mapping [][]itemMapping) map[int]int{
    currMap := seedList
    for i, itemnMap := range mapping { 
        fmt.Println("Running map iteration: ", i)
        currMap = getNextRangeMap(currMap, itemnMap)
    }
    return currMap
}

func part2(input string){
    seedList := getSeedLIst(input) 
    seed2SoilMap := getMap(input, seed2SoildPattern)
    soil2FertMap := getMap(input, soil2FertPattern)
    fert2WaterMap := getMap(input, fert2WaterPattern)
    water2LigtMap := getMap(input, water2LightPattern)
    light2TempMap := getMap(input, light2TempPattern)
    temp2HumidMap := getMap(input, temp2HumidPattern)
    humid2LocationMap := getMap(input, humid2LocationPattern)

    mappings := [][]itemMapping {seed2SoilMap, soil2FertMap, fert2WaterMap, water2LigtMap, light2TempMap, temp2HumidMap, humid2LocationMap}
    seedLocations := getSeedListLocations(seedList, mappings)
    fmt.Println(seedLocations)

    solution := -1
    for minLocation, _ := range seedLocations {
        if solution == -1 || minLocation < solution {
           solution = minLocation 
        }
    }
    fmt.Println(solution)
}

func main(){
    byteData, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    } 
    input := string(byteData)
    //part1(input)
    part2(input)
}
