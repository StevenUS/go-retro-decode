package main

import (
    "fmt"
    "log"
    "strings"
    "time"
    "bufio"
    "os"
)

var ALPHA =  "abcdefghijklmnopqrstuvwxyz"
var CYPHER = "xgczygkhzjklfcwlgbfghvwxyz"
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

    // var text = "zg ghy fcwly wg x gxfk chxckyf, wy fhwhlz byfcwly xcz fxky fhby ghy byghzbyfycgf xby hlzxgyz clyxbly zc ghy gzckyg fw yvybywcy zf wc ghy fxfy lxky wzgh ghy kwxlf xcz hww gw xchzyvy ghyf."
    // text = "gww fxcy zzggybycg gbxcchyf wc zyv fybvyb"
    // text = "cwffhczcxgzwc gygwyyc zczzx xcz hf gyxf xgwhg hlcwfzck gyxghby bylyxfyf xcz ghyzb zylyczycczyf. z hxz gw bywbzgy ghy ycgzby 3bz fhzg gyxghby fzccy z wbwgy zg gwb ghwgy glww ghg jxckzy fwzzgzyz gw hfy byghbczck-hfyb glww. ghzf zylxyyz ghy bylyxfy gy fwby ghxc x wyyk"

    texts := []string{
        "wc gwxbzzck ghy zczzx gyxf wzgh ghy bylyxfy lbwcyff gy jxckzy xcz yzzzy",
        "gyxf. 2 cyw fyfgybf jhfg jwzcyz, 2 fwby cyxg fwczxy",
        "gzckygf xby fwvzck byxlly gxfg",
        "wc-gwxbzzck zwchfycgxgzwc.",
        "glyyg hwfylxky bywwbk lwwkf kbyxg ",
        "fcwly wg ghy gxfk fhwhlz hxvy gyyc cxlghbyz gygwby ghy yckzcyybzck zfllyfycgxgzwc (y.k. vyhzcly fylycgwb by-zyfzkc) ",
        "gww-wyyk bylyxfy cycly gwb zczzx zf cwg hxllyczck xf llxccyz. fxy gy wy fhwhlz byvzfzg xcz fyy zg ghyby zf x gyggyb wxy gw hxvy llxccyz bylyxfyf. ",
        "wzll gy hyllghl zg xll zylyczycczyf wg x gyxghby lzky gxckycz xcz gbwcgycz xby zwcy gbwf yzghyb ghy zczzx gyxf wb ghy hf gyxf cwfllygyly. zg gycwfyf zzggzchlg gw cwllxgwbxgy zc zzggybycg gzfy zwcyf. ",
        "lwcxl xlllzcxgzwc fyghl gwb gx/zyv. ",
        "zg'f kbyxg ghxg wy'by kbwwzck, ghg cxc wy kyg x gbxcflxbycg hzbzck/kbwwgh llxc yvyc zg zg'f jhfg gwb gwzxy xcz fxy chxcky? ",
        "zw wy byxlly hxvy llxcf gw hlkbxzy xzfzc lzky fxfzb fycgzwcyz zhbzck ghy cwflxcy fyygzck? ",
        "xckhlxbjf wzll cw lwckyb gy fhllwbgyz zc 2021. zw wy hxvy xcy llxcf gw hlkbxzy/fzkbxgy gych xll xcz xzfzc xll? ",
        "gxz bwhcz wg zcgybvzywf ",
        "cyyz zyvwlf, ghzlz gwx wc jyckzcf, xcz fwfywcy zyzzcxgyz gw fxzcgxzc jyckzcf jwgf xcz wghyb lbwcyffyf ",
        "cwffhczcxgzwc gygwyyc zczzx xcz hf gyxf xgwhg hlcwfzck gyxghby bylyxfyf xcz ghyzb zylyczycczyf. z hxz gw bywbzgy ghy ycgzby 3bz fhzg gyxghby fzccy z wbwgy zg gwb ghwgy glww ghg jxckzy fwzzgzyz gw hfy byghbczck-hfyb glww. ghzf zylxyyz ghy bylyxfy gy fwby ghxc x wyyk ",
        "xz- lbwzhcg gyxf- fhxby lbwzhcg byghzbyfycgf chycklzfg wzgh ghy yckzcyybzck gyxf xf wyll- ghyzb lhblwfy zf gw ycfhby ghy clxbzgy zf ghyby gygwby fgxbgzck x gzckyg ",
        "xz- kbzfgz- zzfczllzcyz llxcczck fyffzwcf wzgh ghy zczzx gyxf (gzckyg kzck-wggf) ",
        "xz- kbzfgz- lyxbc flbzcgf vzyw zc jzbx gwb zczzx gyxf ",
        "xz- lbwzhcg gyxf (kbzfgz) gw lbzwbzgzzy zyv ghxlzgy wg lzgy gxfk- cwfllygyly xhgwfxgy ghy bylyxfy lbwcyff",
    }

    for _, text := range texts {
        decypher(text)
        fmt.Println("**************************************************************************")
    }

    // end of original js logic
    fmt.Println(time.Since(start))
}

func decypher(text string) {

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

    // end of original js logic
    // fmt.Println(time.Since(start))
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
    // fmt.Println("letters: ", lettersList)
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
    // fmt.Println("words: ", words)
    return words
}


