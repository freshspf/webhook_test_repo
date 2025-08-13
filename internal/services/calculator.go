package services

// FindNarcissisticNumbers returns all narcissistic numbers between 1 and 500
func FindNarcissisticNumbers() []int {
	var result []int
	for i := 1; i <= 500; i++ {
		if isNarcissistic(i) {
			result = append(result, i)
		}
	}
	return result
}

// isNarcissistic checks if a number is narcissistic (sum of each digit cubed equals the number)
func isNarcissistic(n int) bool {
	sum := 0
	original := n

	for n > 0 {
		digit := n % 10
		sum += digit * digit * digit
		n /= 10
	}

	return sum == original
}