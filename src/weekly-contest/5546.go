package main

import "fmt"

func slowestKey(releaseTimes []int, keysPressed string) byte {
    if len(releaseTimes) == 1 {
        return keysPressed[0]
    }
    res := keysPressed[0]
    last, large, time := 0, 0, 0
    for i := range releaseTimes {
        if i > 0 && keysPressed[i-1] == keysPressed[i] {
            time += releaseTimes[i] - last
        } else {
            time = releaseTimes[i] - last
        }
        if time > large {
            large = time
            res = keysPressed[i]
        } else if time == large && keysPressed[i] > res {
            large = time
            res = keysPressed[i]
        }
        last = releaseTimes[i]
    }
    return res
}

func main() {
    a := []int{9,29,49,50}
    fmt.Println(slowestKey(a, "cbcd"))
}
