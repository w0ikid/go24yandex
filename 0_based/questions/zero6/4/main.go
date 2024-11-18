package main
func SumOfArray(array [6]int) int {
    sum := 0
    for _, value := range array {
        sum += value
    }
    return sum
}
