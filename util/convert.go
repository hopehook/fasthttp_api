package util

import (
	"strconv"
)

// map
func Map2KeysValues(m map[string]interface{}) ([]string, []interface{}) {
	keys, values := []string{}, []interface{}{}
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

// slice
func SliceString2Int64(slice []string) []int64 {
	sliceCopy := make([]int64, len(slice))
	for i, j := range slice {
		value, err := strconv.ParseInt(j, 10, 64)
		if err != nil {
			return []int64{}
		}
		sliceCopy[i] = value
	}
	return sliceCopy
}

func SliceByteInterface2String(slice []interface{}) []string {
	sliceCopy := make([]string, len(slice))
	for i, j := range slice {
		value, ok := j.([]byte)
		if !ok {
			return []string{}
		}
		sliceCopy[i] = string(value)
	}
	return sliceCopy
}
