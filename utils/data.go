package utils

func DeleteSliceElement(slice []int64, index int) []int64 {
	return append(slice[:index], slice[index+1:]...)
}
