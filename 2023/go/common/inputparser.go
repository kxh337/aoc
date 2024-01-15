package common

import (
    "os"
    "log"
    "strings"
)

func LoadFile(relpath string) string {
    byteData, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    } 
    input := string(byteData)
    return input
}

func LoadFileLines(relpath string) []string {
    filestr := LoadFile(relpath)
    return  strings.Split(filestr, "\n")
}
