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
    lines := strings.Split(filestr, "\n")
    var res []string
    for _, line := range lines {
        if len(line) != 0{
            res = append(res, line)
        }
    }
    return res
}
