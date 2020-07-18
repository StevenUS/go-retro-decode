package main

import (
    "fmt"
    "strings"
    "time"
)

var ALPHA =  "abcdefghijklmnopqrstuvwxyz";
var CYPHER = "xgczygkhzjklfcwlgbfghvwxyz";

type PossibleWord struct {
    idx int
    possibleLetters []string
}

func main() {
    start := time.Now()
    var text = "zf ghyby x wxy gw hxvy x fzckly fcbzlg ghxg cwcgzkhbyf ywhb lwcxl gxckycz? (fybvyb.xfl, ycv vxbf, lwcxl.lbwlybgzyf, wbycch.lbwlybgzyf ygc.)"

    var words []string
    for _, word := range strings.Fields(text) {
        words = append(words, word)
    }

    // [
    //     [
    //         [d i z]
    //         [m s]
    //     ] ...
    possibleWords := make([][][]string, 0)
    for _, word := range words {
       possibleWords = append(possibleWords, getPossibleWords(word))
    }
    fmt.Println(possibleWords)

    permutations := make([][]string, 0)
    for _, lettersList := range possibleWords {
        permutations = append(permutations, getPermutations(lettersList))
    }

    fmt.Println(permutations)

    fmt.Println(time.Since(start))
}

// accepts a slice of slices of letters which represents a possible word
// the 0 index of the main slice is all of the possible characters of the first
// letter of the word
func getPermutations(lettersList [][]string) []string  {
    var words []string
    for i, letters := range lettersList {
        if (i == 0) {
            words = letters
        } else {
            newWords := make([]string, 0)
            for _, letter := range letters {
                for _, word := range words {
                    newWords = append(newWords, word + letter)
                }
            }
            words = newWords
        }
    }
    return words
}

func getLetters(letter string) []string {
    var letters []string
    for i, c := range CYPHER {
        if string(c) == letter {
            letters = append(letters, string(ALPHA[i]))
        }
    }
    return letters
}

func getPossibleWords(cypheredWord string) [][]string {
    possibleWords := make([][]string, 0)
    for curIdx, letter := range cypheredWord {
        letters := getLetters(string(letter))
        for _, letter := range letters {
            if len(possibleWords) > curIdx {
                possibleWords[curIdx] = append(possibleWords[curIdx], letter)

            } else {
                possibleWords = append(possibleWords, []string{letter})
            }
        }
    }
    if false {
        for key, value := range possibleWords {
            fmt.Println("Key:", key, "Value:", value)
        }
    }
    return possibleWords
}


