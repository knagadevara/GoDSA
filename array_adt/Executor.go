package arrayadt

import (
	"fmt"
	"slices"
)

func Exec() {
	IntArr := MakeArrAdt[int](11)
	fmt.Println(IntArr.RmPop(0))
	IntArr.Addpend(10)
	IntArr.Addpend(20)
	IntArr.Addpend(30)
	IntArr.Addpend(40)
	IntArr.Insert(10, 100)
	IntArr.Insert(0, 99)
	IntArr.Insert(3, 56)
	IntArr.Insert(4, 86)
	IntArr.Insert(9, 16)
	IntArr.Insert(9, 65)
	slices.Sort(*IntArr.adt)
	fmt.Println(IntArr.SortedInsert(39))
	fmt.Println(IntArr.IsSorted())
	fmt.Println(IntArr.ArContainsBS(0))
	fmt.Println(IntArr.ArContainsBS(31))
	fmt.Println("Max: ", Max(IntArr.adt))
	fmt.Println("Min: ", Min(IntArr.adt))
	fmt.Println("Sum: ", Sum(*IntArr.adt))
	IntArr.Display()
	fmt.Println("Concatinated")
	IntArr.Reverse2()
	IntArr.Display()
	var a1 = make([]int, 7)
	a1 = []int{2, 4, 5, 8, 10, 12, 14, 15, 21, 23, 56, 77, 84}
	var a2 = make([]int, 5)
	a2 = []int{1, 3, 5, 7, 9, 11, 13, 15, 16, 17, 36}
	var a3 = make([]int, 7)
	a3 = []int{2, 4, 5, 6, 7, 8, 9}
	var a4 = make([]int, 5)
	a4 = []int{1, 2, 3, 5, 9}
	mergedarr := UnionMergeSA[int](a1, a2)
	for i, v := range mergedarr {
		fmt.Println(i, v)
	}
	IntArr.Concat(a1)
	IntArr.Display()
	intrMrg := IntersectionMergeSA[int](a3, a4)
	for i, v := range intrMrg {
		fmt.Println(i, v)
	}
	diffMSA := DifferenceMergeSA[int](a3, a4)
	for i, v := range diffMSA {
		fmt.Println(i, v)
	}
	a5 := []int{1, 2, 5, 6, 8, 9}
	FindMisngNn(a5)
	a6 := []int{2, 4, 6, 10, 12, 14, 16, 20}
	fmt.Println(FindMissingEO(a6))
}
