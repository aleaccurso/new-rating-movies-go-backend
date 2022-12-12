package utils

func IndexOfInt32(slice []int32, number int32) int {

	for idx, element := range slice {
		if element == number {
			return idx
		}
	}
	return -1
}
