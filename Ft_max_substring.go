package main

import "fmt"

func Ft_max_substring(s string) int {
    seen := make(map[rune]int)

    maxLength := 0
    start := 0

    for i, c := range s {
        if pos, ok := seen[c]; ok && pos >= start {
            start = pos + 1
        }

        currentLength := i - start + 1
        if currentLength > maxLength {
            maxLength = currentLength
        }

        seen[c] = i
    }

    return maxLength
}
