package main

import (
    "fmt"
    "log"
    "strings"
    "time"
    "bufio"
    "os"
)

var ALPHA =  "abcdefghijklmnopqrstuvwxyz";
var CYPHER = "xgczygkhzjklfcwlgbfghvwxyz";
var ENGLISH_WORDS = make(map[string]bool)

func getEnglishWordsFromFile() map[string]bool {
    start := time.Now()

    words := make(map[string]bool)

    file, err := os.Open("./words_alpha.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words[scanner.Text()] = true
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("upload words to map: ", time.Since(start))

    return words

}

func main() {
    ENGLISH_WORDS = getEnglishWordsFromFile()

    start := time.Now()

    var text = "zg ghy fcwly wg x gxfk chxckyf, wy fhwhlz byfcwly xcz fxky fhby ghy byghzbyfycgf xby hlzxgyz clyxbly zc ghy gzckyg fw yvybywcy zf wc ghy fxfy lxky wzgh ghy kwxlf xcz hww gw xchzyvy ghyf."

    var words []string
    for _, word := range strings.Fields(text) {
        words = append(words, word)
    }

    possibleWords := make([][]string, 0)
    for _, word := range words {
       possibleWords = append(possibleWords, getPossibleWords(word))
    }

    permutationLists := make([][]string, 0)
    for _, lettersList := range possibleWords {
        permutationLists = append(permutationLists, getPermutations(lettersList))
    }

    // print results
    for _, list := range permutationLists {
        fmt.Println(list)
    }

    fmt.Println(time.Since(start))
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

func getPossibleWords(cypheredWord string) []string {
    possibleWords := make([]string, 0)
    for curIdx, letter := range cypheredWord {
        letters := getLetters(string(letter))
        for _, decypheredLetter := range letters {
            if len(possibleWords) > curIdx {
                possibleWords[curIdx] = possibleWords[curIdx] + decypheredLetter

            } else {
                possibleWords = append(possibleWords, decypheredLetter)
            }
        }
    }
    return possibleWords
}

// accepts a slice of letters which represents a possible word,
// the 0 index of the slice is all of the possible characters of the first
// letter of the word
func getPermutations(lettersList []string) []string  {
    var words []string
    for i, letters := range lettersList {
        if (i == 0) {
            firstLetters := strings.Split(letters, "")
            words = firstLetters
        } else {
            newWords := make([]string, 0)
            for _, letter := range letters {
                for _, word := range words {
                    newWord := word + string(letter)
                    if len(newWord) == len(lettersList) {
                        if ENGLISH_WORDS[newWord] {
                            newWords = append(newWords, newWord)
                        }
                    } else {
                        newWords = append(newWords, newWord)
                    }
                }
            }
            words = newWords
        }
    }
    return words
}


