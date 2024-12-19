package arrayadt

type AllowedType interface {
	int | string | rune
}

// type
func Sum(arr []int) int {
	var cmpOp int
	for i := 0; i <= len(arr)-1; i++ {
		cmpOp += arr[i]
	}
	return cmpOp
}

func Avg(arr []int) int {
	var cmpOp int
	for i := 0; i <= len(arr)-1; i++ {
		cmpOp += arr[i]
	}
	return cmpOp / len(arr)
}
