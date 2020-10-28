package main

func uniqueOccurrences(arr []int) bool {
    // map(int, int)
    counters := map[int]int{}
    for _, v := range arr {
        counters[v]++
    }

    // set(int)
    times := map[int]struct{}{}
    for _, c := range counters {
        times[c] = struct{}{}
    }
    return len(times) == len(counters)
}
