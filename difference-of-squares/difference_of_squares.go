package diffsquares

func SquareOfSum(n int) int {
	sum := 0
	if n%2 == 0 {
		sum += (n + 1) * (n / 2)
	} else {
		sum += (n+1)*(((n+1)/2)-1) + int(n/2) + 1
	}

	return sum * sum

}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
