package main

func CalculateDigitalRoot(n int) int{
	if n < 10 {
		return n
	}
	sum := 0
	for n > 0 {
		sum = sum + (n%10)
		n = n / 10
	}
	return CalculateDigitalRoot(sum)
}

func main(){

}