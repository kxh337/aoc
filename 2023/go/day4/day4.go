/*
Each card has 2 list of numbers separated by a "|":
    - a list of winning numbers
    - a list of numbers you have

Need to figure out which numbers you have appear appear in the list of wining numbers
1st match makes the card worth one point and each match after the first doubles the point value of that card

part 2 
Win more scratchcards
if card 10 has 5 matching numbers, you win one copy of the cards 11-15
*/

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

type card struct {
    winNums []int
    ownNums []int
}

const numStart = 10
const lhsNumberIndex = 10
const cardPattern = "(\\d+)"

func parseCards(rawCards []string) []card{
    var cards []card
    re := regexp.MustCompile(cardPattern)
    for _, rawCard := range(rawCards) {
        var currCard card
        matches := re.FindAllString(rawCard[numStart:], -1)
        for index, match := range matches {
            num, err := strconv.Atoi(match)
            if err != nil {
                log.Fatal(err)
            }
            if index < lhsNumberIndex {
                currCard.winNums = append(currCard.winNums, num)
            } else {
                currCard.ownNums = append(currCard.ownNums, num)
            }
        }
        cards = append(cards, currCard)
    }
    return cards
}

func calTotalPoints(cards []card)int{
    points := 0
    for _, card := range(cards){
        points += calCardPoints(card)
    }
    return points
}

func calCardPoints(card card)int{
    currCardPoint := 0
     for _, ownNum := range(card.ownNums) {
        if slices.Contains(card.winNums, ownNum){
            if currCardPoint == 0 {
                currCardPoint++
            } else {
                currCardPoint = currCardPoint * 2
            }
        }
    }
    return currCardPoint
}

func calCardMatchs(card card)int{
    cardMatchCount := 0
    fmt.Println(card.winNums, card.ownNums)
     for _, ownNum := range(card.ownNums) {
        if slices.Contains(card.winNums, ownNum){
            cardMatchCount++
        }
    }
    return cardMatchCount
}

func part1(input []string) int{
    cards := parseCards(input)
    return calTotalPoints(cards)
}

func createCardCopies(cardCopyMap map[int] []int, cardCountList []int) {
    for index := 0; index < len(cardCountList); index++ {
        if cardCopyMap[index] != nil{
            for _, cardIndex := range(cardCopyMap[index]) {
                cardCountList[cardIndex] += cardCountList[index]
            }
        }
    }
}

func calCardCopies(cards []card) int{
    cardCopyMap := make(map[int] []int)
    for index, currCard := range(cards) {   
        matches := calCardMatchs(currCard)
        if matches > 0 {
            for cardIndex := index + 1; cardIndex <= (index + matches) && cardIndex < (len(cards)); cardIndex++{
               cardCopyMap[index] = append(cardCopyMap[index], cardIndex)
            }
            //fmt.Println(index, points, cardCopyMap[index])
        }
    }

    cardCountList := make([]int, len(cards))
    for index := 0; index < len(cardCountList); index++ {
        cardCountList[index] = 1
    }

    createCardCopies(cardCopyMap, cardCountList)
    fmt.Println(cardCountList)

    totalCards := 0
    for index := 0; index < len(cardCountList); index++ {
        totalCards += cardCountList[index]
    }

    return totalCards
}

func part2(input []string) int{
    cards := parseCards(input)
    return calCardCopies(cards)
}

func main(){
    byteData, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    } 
    input := string(byteData)
    rawCards := strings.Split(input, "\n")
    rawCards = rawCards[0:len(rawCards)-1]

    //points := part1(rawCards)
    //fmt.Println("Points: ", points)
    totalCards := part2(rawCards)
    fmt.Println("Total Cards: ", totalCards)

}
