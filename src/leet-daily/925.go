package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
    n, t := 0, 0
    for n < len(name) || t < len(typed) {
        if t >= len(typed) {
            return false
        }
        if n < len(name) && name[n] == typed[t] {
            n ++
            t ++
        } else if n > 0 && name[n-1] == typed[t] {
            t ++
        } else {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(isLongPressedName("alex", "aaleex"))
    fmt.Println(isLongPressedName("saeed", "ssaaedd"))
    fmt.Println(isLongPressedName("leelee", "lleeelee"))
    fmt.Println(isLongPressedName("laiden", "laiden"))
    fmt.Println(isLongPressedName("pyplrz", "ppyypllr"))
}
