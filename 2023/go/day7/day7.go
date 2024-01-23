package main

import (
	"aoc/common"
	"fmt"
	"strconv"
	"log"
	"regexp"
)

/*
Playing Camel Cards in the desert
get a list of hands and the goal is to order them based on strength
hand consists of 5 cards.  Ace - 2  where "T" is 10
every hand is exactly one type from strongest to weakest
- five of a kind
- four of a kind
- full house
- 3 of a kind
- 2 pairs
- 1 pair
- high card, i.e. no matching pairs

if the card type is the same, then compare the first card values in both hands
If those are the same move on to comparing the next card in each hand until they differ
*/

type CardRank int8

const (
    HighCard CardRank = iota
    OnePair
    TwoPair
    ThreeKind
    FullHouse
    FourKind
    FiveKind 
)

type hand struct {
    cardstr string
    cardval []int
    kind CardRank
    bid int
}

var cardcharvalmap = map[string]int {
    "A": 14,
    "K": 13,
    "Q": 12,
    "J": 11,
    "T": 10,
    "9": 9,
    "8": 8,
    "7": 7,
    "6": 6,
    "5": 5,
    "4": 4,
    "3": 3,
    "2": 2,
}
const tokenpattern = "\\S+"

func determineRank(currhand hand) CardRank{
    rank := HighCard
    cardmap := make(map [int]int)
    for _, val := range currhand.cardval {
        _, ok := cardmap[val] 
        if !ok {
            cardmap[val] = 1
        } else {
            cardmap[val] = cardmap[val] + 1
        }
    }

    for _, count := range cardmap{
        var nextRank CardRank
        switch count {
            case 2: // pair
                if (rank == ThreeKind){
                    nextRank = FullHouse
                } else if rank == OnePair {
                    nextRank = TwoPair
                } else {
                    nextRank = OnePair
                }
            case 3: // 3 of a kind
                if (rank == OnePair){
                    nextRank = FullHouse
                } else {
                    nextRank = ThreeKind
                }
            case 4:
                rank = FourKind
            case 5:
                rank = FiveKind
        }
        if nextRank > rank {
            rank = nextRank
        }
    }

    return rank
}

func parseHands(lines []string)[]hand{
    var hands [] hand
    re := regexp.MustCompile(tokenpattern)
    for _, line := range lines {
        var currentHand hand
        matches := re.FindAllString(line, -1)
        if (len(matches) == 2){
            currentHand.cardstr = matches[0]
            bid, err := strconv.Atoi(matches[1])
            if err != nil{
                log.Fatal(err)
            }
            currentHand.bid = bid
            var cardvals [] int
            for _, lit := range currentHand.cardstr {
                cardvals = append(cardvals, cardcharvalmap[string(lit)])
            }
            currentHand.cardval = cardvals
            currentHand.kind = determineRank(currentHand)
            hands = append(hands, currentHand)
        }
    }
    return hands
}

func rankHands(hands []hand)[]hand{
    // Just do insertion sort here 
    for index := 0; index < len(hands); index ++ {
        var nextHand =  hands[index]
        var nextIndex =  index
        for nextSubIndex, hand := range hands[index+1:]{
            if hand.kind < nextHand.kind{
                nextHand = hand 
                nextIndex = nextSubIndex + index + 1
            } else if hand.kind == nextHand.kind{
                // check for lowest card value in sequence
                for cardIndex := 0; cardIndex < len(hand.cardval); cardIndex++ {
                    if hand.cardval[cardIndex] < nextHand.cardval[cardIndex] {
                        nextHand = hand 
                        nextIndex = nextSubIndex + index + 1
                        break
                    } else if hand.cardval[cardIndex] > nextHand.cardval[cardIndex]{
                        break
                    }
                }
            }
        }
        // swap 
        if index != nextIndex {
            tmpHand := hands[index]
            hands[index] = nextHand
            hands[nextIndex] = tmpHand
        }
        
    }
    return hands
}

func main() {
    lines := common.LoadFileLines("input.txt") 
    hands := parseHands(lines)
    rankedHands := rankHands(hands)

    sum := 0
    for rank, hand := range rankedHands{
        fmt.Println(hand.kind, hand.cardval, hand.bid, rank)
        sum = sum + (rank + 1) * hand.bid
    }

    fmt.Println(sum)
}
