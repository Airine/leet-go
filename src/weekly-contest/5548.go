package main

import "math"

func minEffort(heights [][]int, x int, y int) int {
    res := 10 ^ 6
    if x - 1 == 0 && y == 0 {
        diff := int(math.Abs(float64(heights[x][y]-heights[x-1][y])))
        return diff
    } else if x == 0 && y - 1 == 0 {
        diff := int(math.Abs(float64(heights[x][y]-heights[x][y-1])))
        return diff
    }
    if x > 0 {
        tmpMax := minEffort(heights, x-1, y)
        diff := int(math.Abs(float64(heights[x][y]-heights[x-1][y])))
        if diff > tmpMax {
            tmpMax = diff
        }
        if tmpMax < res {
            res = tmpMax
        }
    }
    if x < len(heights)-1 {
        tmpMax := minEffort(heights, x+1, y)
        diff := int(math.Abs(float64(heights[x][y]-heights[x+1][y])))
        if diff > tmpMax {
            tmpMax = diff
        }
        if tmpMax < res {
            res = tmpMax
        }
    }
    if y > 0 {
        tmpMax := minEffort(heights, x, y-1)
        diff := int(math.Abs(float64(heights[x][y]-heights[x][y-1])))
        if diff > tmpMax {
            tmpMax = diff
        }
        if tmpMax < res {
            res = tmpMax
        }
    }
    if y < len(heights[0])-1 {
        tmpMax := minEffort(heights, x, y+1)
        diff := int(math.Abs(float64(heights[x][y]-heights[x][y+1])))
        if diff > tmpMax {
            tmpMax = diff
        }
        if tmpMax < res {
            res = tmpMax
        }
    }
    return res
}

func minimumEffortPath(heights [][]int) int {
    row, col := len(heights), len(heights[0])

    scores := make([][]int, row)
    for i := range scores {
        scores[i] = make([]int, col)
    }
    for i := 0; i < row; i ++ {
        for j := 0; j < col; j ++ {
            scores[i][j] = 10^6
        }
    }
    return minEffort(scores, row-1, col-1)
}

func main() {
    
}
