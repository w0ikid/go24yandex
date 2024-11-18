package main

func FindMaxKey(m map[int]int) int {
    maxKey := -1
    for key := range m {
        if key > maxKey {
            maxKey = key
        }
    }
    return maxKey
}
