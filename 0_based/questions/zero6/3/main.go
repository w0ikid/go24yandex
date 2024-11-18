package main

func FindMinMaxInArray(array [10]int) (int, int) {
    min, max := array[0], array[0]

    for _, value := range array {
        if value < min {
            min = value
        }
        if value > max {
            max = value
        }
    }
    return min, max
}