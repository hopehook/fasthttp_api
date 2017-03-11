package util

func MaxOfSlice(slice []int64) (index int64, value int64) {
	for i, j := range slice {
		if value < j {
			index = int64(i)
			value = int64(j)
		}
	}
	return
}
