package main

import (
    "fmt"
    "math"
    "sort"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
    //res := make([][]int, R*C)
    var res [][]int
    for i := 0; i < R; i++ {
        for j := 0; j < C; j++ {
            res = append(res, []int{i, j})
        }
    }
    sort.Slice(res, func(i, j int) bool {
        r1, c1 := res[i][0], res[i][1]
        r2, c2 := res[j][0], res[j][1]
        return math.Abs(float64(r1-r0)) + math.Abs(float64(c1-c0)) < math.Abs(float64(r2-r0)) + math.Abs(float64(c2-c0))
    })
    return res
}

func main() {
    fmt.Println(allCellsDistOrder(2, 3, 1, 2))
}
