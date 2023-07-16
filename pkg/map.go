package util

func Map(array []int, f func(int) int) []int {
	var result []int
	for _, v := range array {
		result = append(result, f(v))
	}
	return result
}
