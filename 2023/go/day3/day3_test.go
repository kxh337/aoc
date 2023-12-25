package main

import (
    "testing"
)

func TestBasicGear(t *testing.T){
    input := []string{"353*353"}
    sum := Part2(input)
    if sum != (353 * 353) {
        t.Fatalf("Sum: %d is not equal to expected value of :%d", sum, (353 * 353))
    }
}

func TestEdge(t *testing.T){
    input := []string{"353*....", 
                      "....353."}
    sum := Part2(input)
    if sum != (353 * 353) {
        t.Fatalf("Sum: %d is not equal to expected value of :%d", sum, (353 * 353))
    }
}

func TestFindNum(t *testing.T){
    input := "....353."
    nums := findNumbers(input, len(input) - 1)
    if (nums[0].number != 353 ) {
        t.Fatalf("Found number: %d instead of 353", nums[0].number)
    }
}
