package arrayadt

import (
	"cmp"
	"fmt"
)

type AllowedType interface {
	int | string | rune
}

func Max(arr *[]int) int {
	var cmpOp int
	for i := 0; i <= len(*arr)-1; i++ {
		if (*arr)[i] > cmpOp {
			cmpOp = (*arr)[i]
		}
	}
	return cmpOp
}

func Min(arr *[]int) int {
	var cmpOp int
	for i := 0; i <= len(*arr)-1; i++ {
		if (*arr)[i] < cmpOp {
			cmpOp = (*arr)[i]
		}
	}
	return cmpOp
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

func UnionMergeSA[T cmp.Ordered](A1, A2 []T) []T {
	var i, j, k, A1Len, A2Len, A3Len int
	A1Len = len(A1)
	A2Len = len(A2)
	var A3 = make([]T, (A1Len + A2Len))
	A3Len = len(A3)
	for {
		if i >= A3Len {
			break
		} else if j == A1Len && k < A2Len {
			for _, kv := range A2[k:A2Len] {
				A3[i] = kv
				i++
			}
			break
		} else if k == A2Len && j < A1Len {
			for _, jv := range A1[j:A1Len] {
				A3[i] = jv
				i++
			}
			break
		} else if j != A1Len && k != A2Len {
			if A1[j] < A2[k] {
				A3[i] = A1[j]
				if j < A1Len {
					j++
				}
			} else if A1[j] > A2[k] {
				A3[i] = A2[k]
				if k < A2Len {
					k++
				}
			} else if A1[j] == A2[k] {
				A3[i] = A2[k]
				if k < A2Len && j < A1Len {
					k++
					j++
				}
			}
			i++
		} else {
			break
		}
	}
	return A3
}

func IntersectionMergeSA[T cmp.Ordered](A1, A2 []T) []T {
	var i, j, k, A1Len, A2Len int
	A1Len = len(A1) - 1
	A2Len = len(A2) - 1
	var A3 = make([]T, (A1Len + A2Len))
	for {
		if j >= A1Len && k >= A2Len {
			break
		} else {
			if A1[j] < A2[k] {
				if j < A1Len {
					j++
				}
			}
			if A1[j] > A2[k] {
				if k < A2Len {
					k++
				}
			}
			if A1[j] == A2[k] {
				A3[i] = A2[k]
				if k < A2Len && j < A1Len {
					k++
					j++
					i++
				}
			} else {
				if j == A1Len && k < A2Len {
					k++
				} else if k == A2Len && j < A1Len {
					j++
				}
			}
		}

	}
	return A3
}

func DifferenceMergeSA[T cmp.Ordered](A1, A2 []T) []T {
	var i, j, k, A1Len, A2Len int
	A1Len = len(A1) - 1
	A2Len = len(A2) - 1
	var A3 = make([]T, (A1Len))
	for {

		if j >= A1Len && k >= A2Len {
			break
		}
		if A1[j] == A2[k] {
			if j < A1Len && k < A2Len {
				j++
				k++
				continue
			}
		} else {
			if A1[j] < A2[k] {
				A3[i] = A1[j]
				if j < A1Len {
					j++
					i++
					continue
				}
			} else if A1[j] > A2[k] {
				if k < A2Len {
					k++
					continue
				}
			}
			if j >= A1Len && k < A2Len {
				k++
			}
			if k >= A2Len && j < A1Len {
				j++
			}
		}

	}
	return A3
}

// using mathematical equation n*(n+1)/2
func FindFirstMissingElem(Array []int) int {
	sum := Sum(Array)
	ArrLen := len(Array) - 1
	fml := (Array[ArrLen] * (Array[ArrLen] + 1)) / 2
	return fml - sum
}

// requires A[0] and A[1] to be present
// will not work on contigous missing elements
func FindMissingEO(Array []int) []int {
	var lenArr, diff, i, j int
	var missingElem = make([]int, ((len(Array)-1)/2)-1)
	diff = Array[1] - Array[0]
	lenArr = len(Array) - 1
	for {
		if i >= lenArr {
			break
		}
		if (Array[i+1] - Array[i]) != diff {
			missingElem[j] = Array[i] + diff
			j++
		}
		i++
	}
	return missingElem
}

// Works only for a sorted array
func FindMisngNn(Array []int) {
	max := Array[len(Array)-1]
	var missingElem = make([]bool, max+1)
	for i := 0; i <= len(Array)-1; i++ {
		missingElem[Array[i]] = true
	}
	for i := Array[0]; i < max; i++ {
		if missingElem[i] {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
