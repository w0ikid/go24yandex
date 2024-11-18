package main

func FiveSteps(array [5]int) [5]int{
	var arrayReverse[5]int
	k := len(array)
	for i := 0; i < len(array); i++ {
		arrayReverse[i] = array[k-1]
		k--
	}
	return arrayReverse
}

func main(){

}