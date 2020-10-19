package main

import "fmt"

func backspaceCompare(S string, T string) bool {
    sPtr, tPtr := len(S)-1, len(T)-1
    sBack, tBack := 0, 0
    for sPtr >= 0 || tPtr >= 0 {
        if sPtr >= 0 {
            if S[sPtr] == '#' {
                sBack++
                sPtr--
                continue
            } else if sBack > 0 {
                sPtr--
                sBack--
                continue
            }
        }
        if tPtr >= 0 {
            if T[tPtr] == '#' {
                tBack++
                tPtr--
                continue
            } else if tBack > 0 {
                tPtr--
                tBack--
                continue
            }
        }

        if sPtr >= 0 && tPtr >= 0 && S[sPtr] != T[tPtr] {
            return false
        }
        sPtr --
        tPtr --
    }
    return sPtr == tPtr
}

func main() {
    fmt.Println(backspaceCompare("xywrrmp", "xywrrmu#p"))
    fmt.Println(backspaceCompare("ab#c", "ad#c"))
    fmt.Println(backspaceCompare("ab##", "c#d#"))
    fmt.Println(backspaceCompare("a##c", "#a#c"))
    fmt.Println(backspaceCompare("a#c", "b"))
    fmt.Println(backspaceCompare("bxj##tw", "bxj###tw"))
}