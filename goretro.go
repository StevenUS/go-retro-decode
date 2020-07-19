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

type PossibleWord struct {
    idx int
    possibleLetters []string
}

type WordSet map[string]struct{}

func getEnglishWords() map[string]bool {
    start := time.Now()

    words := make(map[string]bool)

    if words["test"] {
        fmt.Println("success")
    }

    file, err := os.Open("../go-words/english-words/words_alpha.txt")
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
    englishWords := getEnglishWords()
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

    permutationLists := make([][]string, 0)
    for _, lettersList := range possibleWords {
        permutationLists = append(permutationLists, getPermutations(lettersList))
    }

    // fmt.Println(permutationLists)

    getAcutalWords(englishWords, permutationLists)

    fmt.Println(time.Since(start))
}

func getAcutalWords(englishWords map[string]bool, permutationLists [][]string) {

    actualWords := make(map[int][]string)
    for i, permutations := range permutationLists {
        for _, permutation := range permutations {

            if (englishWords[permutation]) {
                if _, ok := actualWords[i]; ok {
                    actualWords[i] = append(actualWords[i], string(permutation))
                } else {
                    actualWords[i] = []string{string(permutation)}
                }

            }

        }
    }

    if true {
        for key, value := range actualWords {
            fmt.Println("Key:", key, "Value:", value)
        }
    }
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


