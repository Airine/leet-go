package main

func wordBreak(s string, wordDict []string) bool {
    wordSet := map[string]struct{}{}
    for _, v := range wordDict {
        wordSet[v] = struct{}{}
    }
    dp := make([]bool, len(s)+1)
    dp[0] = true
    for i := range s {
        for j := i+1; j < len(s)+1; j++ {
            if dp[i] {
                if _, ok := wordSet[s[i:j]]; ok {
                    dp[j] = true
                }
            }
        }
    }
    return dp[len(s)]
}
