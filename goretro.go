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

    possibleWords := make([]map[int][]string, 0)
    for _, word := range words {
       possibleWords = append(possibleWords, getPossibleWords(word))
    }

    for _, possibleWord := range possibleWords {

        // [['t', 'x']
        // ['t', 'h', 'a']
        // ['e', 'n']]
        listOfLettersLists := make([][]string, 0, len(possibleWord))

        for _, letters := range possibleWord {

            listOfLettersLists = append(listOfLettersLists, letters)
        }

        fmt.Println(listOfLettersLists)
        // for _, letterList := range listOfLettersLists {
        //     getPermutations(letterList)
        // }


    }

     fmt.Println(time.Since(start))
}

func getAlphaIndicies(letter string) []string {
    var letters []string
    for i, c := range CYPHER {
        if string(c) == letter {
            letters = append(letters, getLetter(i))
        }
    }
    return letters
}

func getLetter(idx int) string {
    return string(ALPHA[idx])
}



// ex output: "is there"
// { '0': [ 'd', 'i', 'z' ], '1': [ 'm', 's' ] }
// {
//   '0': [ 'b', 'f', 'q', 't' ],
//   '1': [ 'h', 'u' ],
//   '2': [ 'e', 'y' ],
//   '3': [ 'r' ],
//   '4': [ 'e', 'y' ]
// }
func getPossibleWords(cypheredWord string) map[int][]string {
    possibleWords := make(map[int][]string)
    for curIdx, letter := range cypheredWord {
        letters := getAlphaIndicies(string(letter))
        for _, letter := range letters {
            if val, ok := possibleWords[curIdx]; ok {
                possibleWords[curIdx] = append(val, letter)

            } else {
                possibleWords[curIdx] = []string{letter}
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

// func getPermutations(letters [][]string) {
//     var words []string
//     for i, letter := range letters {
//         if i == 0 {
//             words = []string{letter}
//         } else {
//             var newWords = make([]string,0)
//             for _, word := range words {
//                 newWords = append(newWords, word)
//             }
//
//         }
//     }
//     fmt.Println(words)
//
// }

